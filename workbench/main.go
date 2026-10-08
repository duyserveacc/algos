package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
PASTE THE COMPLETE PROBLEM BELOW THIS LINE

Definition for singly-linked list:

type ListNode struct {
    Val  int
    Next *ListNode
}

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list. You may assume the two numbers do not contain any leading zero, except the number 0 itself.

In the examples, each linked list is serialized as an array of node values.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

PASTE THE COMPLETE PROBLEM ABOVE THIS LINE
===============================================================================
*/

// =============================================================================
// PROBLEM-SPECIFIC TYPES AND ADAPTERS

// ListNode is a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// parseAddTwoNumbersInput converts this problem's serialized list arguments
// into the pointer-based values expected by the judge function.
func parseAddTwoNumbersInput(input string) (*ListNode, *ListNode) {
	var firstValues []int
	var secondValues []int
	if err := testrunner.DecodeJSONArgs(input, &firstValues, &secondValues); err != nil {
		panic(fmt.Errorf("parse add two numbers input: %w", err))
	}

	return linkedListFromValues(firstValues), linkedListFromValues(secondValues)
}

// linkedListFromValues builds a new list in the serialized value order.
func linkedListFromValues(values []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, value := range values {
		tail.Next = &ListNode{Val: value}
		tail = tail.Next
	}
	return dummy.Next
}

// formatAddTwoNumbersOutput serializes the returned list in statement form.
func formatAddTwoNumbersOutput(head *ListNode) string {
	values := make([]int, 0)
	visited := make(map[*ListNode]struct{})
	for node := head; node != nil; node = node.Next {
		if _, exists := visited[node]; exists {
			panic("format add two numbers output: linked list contains a cycle")
		}
		visited[node] = struct{}{}
		values = append(values, node.Val)
	}

	encoded, err := json.Marshal(values)
	if err != nil {
		panic(fmt.Errorf("format add two numbers output: %w", err))
	}
	return string(encoded)
}

// =============================================================================

// sampleCases contains the examples that every registered solution must pass.
// Each input is a positional JSON array containing the statement's serialized
// linked-list values. The problem-specific parser converts them to *ListNode.
var sampleCases = []testrunner.Case{
	{
		Name:  "Example 1",
		Input: `[[2,4,3],[5,6,4]]`,
		Want:  `[7,0,8]`,
	},
	{
		Name:  "Example 2",
		Input: `[[0],[0]]`,
		Want:  `[0]`,
	},
	{
		Name:  "Example 3",
		Input: `[[9,9,9,9,9,9,9],[9,9,9,9]]`,
		Want:  `[8,9,9,9,0,0,0,1]`,
	},
}

// archiveReady is set to true only after the problem, samples, three solutions,
// analysis, and regression coverage are complete and verified.
const archiveReady = false

// solutions lists every implementation that the sample runner checks. Keep the
// current attempt registered while developing it. When Codex adds three final
// approaches, every approach is registered here and receives the same cases.
var solutions = []testrunner.Solution{
	{Name: "current attempt", Run: solveCurrentAttempt},
}

// solveCurrentAttempt is the only function you need to change initially. It
// receives the exact sample input string and returns the expected output form.
// Parse with parseAddTwoNumbersInput, keep the algorithm strongly typed, and
// serialize its returned list with formatAddTwoNumbersOutput.
// Leading and trailing whitespace is ignored during comparison; internal
// whitespace remains significant.
func solveCurrentAttempt(input string) string {
	l1, l2 := parseAddTwoNumbersInput(input)
	_, _ = l1, l2
	// Implement your algorithm and return the formatted answer.
	return ""
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct:
- First failing assumption or bottleneck:
- Relevant concept and invariant:

Solution 1 - Direct:
- Correctness:
- Time complexity:
- Auxiliary-space complexity:

Solution 2 - Improved:
- Correctness:
- Time complexity:
- Auxiliary-space complexity:

Solution 3 - Preferred:
- Correctness:
- Time complexity:
- Auxiliary-space complexity:

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
