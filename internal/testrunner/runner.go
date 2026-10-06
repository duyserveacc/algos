// Package testrunner runs text-input/text-output examples against one or more
// algorithm implementations.
package testrunner

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// DecodeJSONArgs decodes a positional JSON array into the provided destinations.
func DecodeJSONArgs(input string, destinations ...any) error {
	var arguments []json.RawMessage
	if err := json.Unmarshal([]byte(input), &arguments); err != nil {
		return fmt.Errorf("decode JSON arguments: %w", err)
	}
	if len(arguments) != len(destinations) {
		return fmt.Errorf("decode JSON arguments: got %d arguments, want %d", len(arguments), len(destinations))
	}

	for index, argument := range arguments {
		if err := json.Unmarshal(argument, destinations[index]); err != nil {
			return fmt.Errorf("decode JSON argument %d: %w", index+1, err)
		}
	}
	return nil
}

// Case describes one sample input and its expected output.
type Case struct {
	Name  string
	Input string
	Want  string
}

// Solution names one implementation that accepts text input and returns text
// output.
type Solution struct {
	Name string
	Run  func(string) string
}

// Result records one solution/case comparison.
type Result struct {
	Solution string
	Sample   string
	Input    string
	Want     string
	Got      string
	Passed   bool
}

// Run evaluates every solution against every case and writes a concise report.
func Run(output io.Writer, cases []Case, solutions []Solution) error {
	if len(cases) == 0 {
		fmt.Fprintln(output, "No sample cases configured.")
		return nil
	}
	if len(solutions) == 0 {
		return fmt.Errorf("no solutions are registered")
	}
	if err := validate(cases, solutions); err != nil {
		return err
	}

	failures := 0
	for _, result := range Evaluate(cases, solutions) {
		if result.Passed {
			fmt.Fprintf(output, "PASS  %s / %s\n", result.Solution, result.Sample)
			continue
		}

		failures++
		fmt.Fprintf(output, "FAIL  %s / %s\n", result.Solution, result.Sample)
		fmt.Fprintf(output, "  input: %q\n", result.Input)
		fmt.Fprintf(output, "  want:  %q\n", result.Want)
		fmt.Fprintf(output, "  got:   %q\n", result.Got)
	}

	if failures > 0 {
		return fmt.Errorf("%d sample checks failed", failures)
	}
	return nil
}

// Evaluate returns every solution/case comparison without printing a report.
func Evaluate(cases []Case, solutions []Solution) []Result {
	results := make([]Result, 0, len(cases)*len(solutions))
	for _, solution := range solutions {
		for _, sample := range cases {
			got := solution.Run(sample.Input)
			results = append(results, Result{
				Solution: solution.Name,
				Sample:   sample.Name,
				Input:    sample.Input,
				Want:     sample.Want,
				Got:      got,
				Passed:   normalizeOutput(got) == normalizeOutput(sample.Want),
			})
		}
	}
	return results
}

func validate(cases []Case, solutions []Solution) error {
	for index, sample := range cases {
		if strings.TrimSpace(sample.Name) == "" {
			return fmt.Errorf("sample %d has an empty name", index+1)
		}
	}

	names := make(map[string]struct{}, len(solutions))
	for index, solution := range solutions {
		if strings.TrimSpace(solution.Name) == "" {
			return fmt.Errorf("solution %d has an empty name", index+1)
		}
		if solution.Run == nil {
			return fmt.Errorf("solution %q has no function", solution.Name)
		}
		if _, exists := names[solution.Name]; exists {
			return fmt.Errorf("solution name %q is registered more than once", solution.Name)
		}
		names[solution.Name] = struct{}{}
	}
	return nil
}

func normalizeOutput(value string) string {
	return strings.TrimSpace(strings.ReplaceAll(value, "\r\n", "\n"))
}
