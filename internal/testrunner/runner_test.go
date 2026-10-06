package testrunner

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	type input struct {
		Nums   []int `json:"nums"`
		Target int   `json:"target"`
	}

	got, err := DecodeJSON[input](`{"nums":[2,7,11,15],"target":9}`)
	if err != nil {
		t.Fatalf("DecodeJSON() error = %v, want nil", err)
	}
	want := input{Nums: []int{2, 7, 11, 15}, Target: 9}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("DecodeJSON() = %#v, want %#v", got, want)
	}
}

func TestDecodeJSONReturnsContextualError(t *testing.T) {
	_, err := DecodeJSON[map[string]int](`{"target":"nine"}`)
	if err == nil || !strings.Contains(err.Error(), "decode JSON input") {
		t.Fatalf("DecodeJSON() error = %v, want contextual decoding error", err)
	}
}

func TestRunReportsPassesAndFailures(t *testing.T) {
	cases := []Case{{Name: "Example", Input: "hello", Want: "HELLO"}}
	solutions := []Solution{
		{Name: "correct", Run: strings.ToUpper},
		{Name: "incorrect", Run: func(input string) string { return input }},
	}
	var output bytes.Buffer

	err := Run(&output, cases, solutions)
	if err == nil || err.Error() != "1 sample checks failed" {
		t.Fatalf("Run() error = %v, want one failed check", err)
	}
	if !strings.Contains(output.String(), "PASS  correct / Example") {
		t.Fatalf("Run() output missing pass result:\n%s", output.String())
	}
	if !strings.Contains(output.String(), "FAIL  incorrect / Example") {
		t.Fatalf("Run() output missing failure result:\n%s", output.String())
	}
}

func TestRunRejectsInvalidRegistration(t *testing.T) {
	cases := []Case{{Name: "Example"}}
	solutions := []Solution{
		{Name: "duplicate", Run: strings.ToUpper},
		{Name: "duplicate", Run: strings.ToLower},
	}

	err := Run(&bytes.Buffer{}, cases, solutions)
	if err == nil || !strings.Contains(err.Error(), "registered more than once") {
		t.Fatalf("Run() error = %v, want duplicate-name error", err)
	}
}

func TestRunAllowsNoConfiguredCases(t *testing.T) {
	var output bytes.Buffer

	if err := Run(&output, nil, nil); err != nil {
		t.Fatalf("Run() error = %v, want nil", err)
	}
	if !strings.Contains(output.String(), "No sample cases configured") {
		t.Fatalf("Run() output missing empty-workbench guidance: %s", output.String())
	}
}
