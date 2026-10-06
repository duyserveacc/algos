package main

import (
	"testing"

	"github.com/duyserveacc/algos/internal/testrunner"
)

func TestConfiguredSamplesAgainstEverySolution(t *testing.T) {
	if len(sampleCases) == 0 {
		t.Skip("add workbench samples to enable active-problem checks")
	}

	for _, result := range testrunner.Evaluate(sampleCases, solutions) {
		if !result.Passed {
			t.Errorf("%s / %s: got %q, want %q", result.Solution, result.Sample, result.Got, result.Want)
		}
	}
}
