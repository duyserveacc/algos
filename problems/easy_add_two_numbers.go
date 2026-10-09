//go:build ignore

// Difficulty: Easy

package main

import (
	"encoding/json"
	"fmt"
	"math/big"
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
	{
		Name:  "Single digits without carry",
		Input: `[[1],[2]]`,
		Want:  `[3]`,
	},
	{
		Name:  "Single digits produce a final carry",
		Input: `[[5],[5]]`,
		Want:  `[0,1]`,
	},
	{
		Name:  "First number is zero",
		Input: `[[0],[7,3]]`,
		Want:  `[7,3]`,
	},
	{
		Name:  "Second number is zero",
		Input: `[[7,3],[0]]`,
		Want:  `[7,3]`,
	},
	{
		Name:  "Carry propagates through longer first list",
		Input: `[[9,9,9],[1]]`,
		Want:  `[0,0,0,1]`,
	},
	{
		Name:  "Carry propagates through longer second list",
		Input: `[[1],[9,9,9]]`,
		Want:  `[0,0,0,1]`,
	},
	{
		Name:  "Carry clears before the final digit",
		Input: `[[8,1,2],[7,8,3]]`,
		Want:  `[5,0,6]`,
	},
	{
		Name:  "Maximum length with carry into a new node",
		Input: `[[9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9],[1]]`,
		Want:  `[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]`,
	},
}

// archiveReady is set to true only after the problem, samples, three solutions,
// analysis, and regression coverage are complete and verified.
const archiveReady = true

// solutions lists every implementation that the sample runner checks. Keep the
// current attempt registered while developing it. When Codex adds three final
// approaches, every approach is registered here and receives the same cases.
var solutions = []testrunner.Solution{
	{Name: "direct: arbitrary-precision conversion", Run: solveCurrentAttempt},
	{Name: "improved: recursive carry", Run: solveRecursiveCarry},
	{Name: "preferred: iterative carry", Run: solveIterativeCarry},
}

// solveCurrentAttempt is the only function you need to change initially. It
// receives the exact sample input string and returns the expected output form.
// Parse with parseAddTwoNumbersInput, keep the algorithm strongly typed, and
// serialize its returned list with formatAddTwoNumbersOutput.
// Leading and trailing whitespace is ignored during comparison; internal
// whitespace remains significant.
func solveCurrentAttempt(input string) string {
	l1, l2 := parseAddTwoNumbersInput(input)
	return formatAddTwoNumbersOutput(addTwoNumbersViaBigInt(l1, l2))
}

// addTwoNumbersViaBigInt preserves the current attempt: convert both lists to
// arbitrary-precision integers, add them, and convert the sum back to a list.
func addTwoNumbersViaBigInt(l1, l2 *ListNode) *ListNode {
	listToInteger := func(head *ListNode) *big.Int {
		digits := make([]byte, 0)
		for node := head; node != nil; node = node.Next {
			digits = append(digits, byte('0'+node.Val))
		}
		for left, right := 0, len(digits)-1; left < right; left, right = left+1, right-1 {
			digits[left], digits[right] = digits[right], digits[left]
		}

		value, ok := new(big.Int).SetString(string(digits), 10)
		if !ok {
			panic("convert linked list to integer: invalid digits")
		}
		return value
	}

	sum := new(big.Int).Add(listToInteger(l1), listToInteger(l2))
	digits := sum.String()
	dummy := &ListNode{}
	tail := dummy
	for index := len(digits) - 1; index >= 0; index-- {
		tail.Next = &ListNode{Val: int(digits[index] - '0')}
		tail = tail.Next
	}

	return dummy.Next
}

func solveRecursiveCarry(input string) string {
	l1, l2 := parseAddTwoNumbersInput(input)
	return formatAddTwoNumbersOutput(addTwoNumbersRecursive(l1, l2))
}

// addTwoNumbersRecursive adds one digit per call and carries any overflow into
// the next pair of nodes.
func addTwoNumbersRecursive(l1, l2 *ListNode) *ListNode {
	return addDigitsRecursive(l1, l2, 0)
}

func addDigitsRecursive(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	total := carry
	if l1 != nil {
		total += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		total += l2.Val
		l2 = l2.Next
	}

	return &ListNode{
		Val:  total % 10,
		Next: addDigitsRecursive(l1, l2, total/10),
	}
}

func solveIterativeCarry(input string) string {
	l1, l2 := parseAddTwoNumbersInput(input)
	return formatAddTwoNumbersOutput(addTwoNumbers(l1, l2))
}

// addTwoNumbers is the judge-facing function and delegates to the preferred
// iterative implementation.
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	return addTwoNumbersIterative(l1, l2)
}

// addTwoNumbersIterative walks the reverse-order digits from least to most
// significant while maintaining the carry for the next position.
func addTwoNumbersIterative(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		total := carry
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}

		tail.Next = &ListNode{Val: total % 10}
		tail = tail.Next
		carry = total / 10
	}

	return dummy.Next
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: Converting each reverse-order list to an arbitrary-precision
  integer avoids native integer overflow, produces the correct mathematical
  sum, and converts that sum back to the required reverse-order list.
- First failing assumption or bottleneck: There is no correctness failure for
  the stated constraints. The bottleneck is unnecessary decimal and big-integer
  conversion, which allocates intermediate byte slices, strings, and big.Int
  storage and bypasses the linked-list structure that already exposes digits in
  addition order.
- Relevant concept and invariant: See
  concepts/data-structures/linked-list/README.md. Before processing a position,
  carry is exactly the overflow from the less-significant position. The result
  node is (digit1 + digit2 + carry) % 10, and the next carry is that sum / 10.

Solution 1 - Direct arbitrary-precision conversion:
- Correctness: Reversing each node sequence creates the conventional decimal
  representation of the same number. big.Int addition returns their exact sum,
  and reading the sum string backward restores the required node order.
- Concrete example: For l1 = [2,4,3] and l2 = [5,6,4], listToInteger first
  builds "243" and reverses it to "342", then builds "564" and reverses it to
  "465". SetString parses 342 and 465, and Add produces 807. The final loop
  reads the string "807" at indexes 2, 1, and 0, appending nodes 7, 0, and 8,
  so the returned list is [7,0,8].
- Time complexity: The explicit list and string passes take Theta(n + m + r),
  where r is the result length. Decimal big-integer parsing and formatting are
  implementation-dependent; a conservative schoolbook bound is O(N^2), where
  N = max(n, m). Input values do not provide an asymptotic early exit, so best,
  average, and worst cases all process every input and output digit.
- Auxiliary-space complexity: O(N) excluding the returned list, for digit
  buffers, decimal strings, and big.Int words.

Solution 2 - Improved recursive carry:
- Correctness: Each call emits the correct digit for its position from the two
  available digits and incoming carry, then passes exactly the overflow to the
  next position. The base case stops only when both inputs and the carry are
  exhausted, so induction over the remaining nodes proves the full sum.
- Concrete example: For l1 = [9,9] and l2 = [1], the first call computes
  total = 9 + 1 + 0 = 10, stores digit 0, advances both pointers, and recurses
  with carry 1. The second call computes total = 9 + 0 + 1 = 10, stores another
  0, advances l1 to nil, and recurses with carry 1. The third call has two nil
  pointers, computes total = 1, and stores digit 1. The next call reaches the
  base case; unwinding links the nodes as [0,0,1].
- Time complexity: Theta(N) in the best, average, and worst cases because each
  node and any final carry are processed once.
- Auxiliary-space complexity: O(N) excluding the returned list because the
  recursion keeps one stack frame per output position.

Solution 3 - Preferred iterative carry:
- Correctness: The loop maintains the same carry invariant as the recursive
  approach. It appends one correct result digit per iteration and continues
  until neither input nor carry remains, so the returned list represents the
  complete sum.
- Concrete example: For l1 = [8,1,2] and l2 = [7,8,3], the first iteration has
  carry 0, computes total = 8 + 7 = 15, appends 5, sets carry to 1, and advances
  both pointers. The second computes total = 1 + 8 + 1 = 10, appends 0, and
  keeps carry 1. The third computes total = 2 + 3 + 1 = 6, appends 6, and clears
  carry to 0. Both pointers are now nil, so the loop stops with [5,0,6].
- Time complexity: Theta(N) in the best, average, and worst cases because every
  input node and any final carry are processed once.
- Auxiliary-space complexity: O(1) excluding the returned list; only pointers,
  the current total, and carry are retained.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
