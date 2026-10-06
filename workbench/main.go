package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/duyserveacc/algos/internal/testrunner"
)

/*
===============================================================================
PASTE THE COMPLETE PROBLEM BELOW THIS LINE

You are given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:

2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n^2) time complexity?

PASTE THE COMPLETE PROBLEM ABOVE THIS LINE
===============================================================================
*/

// sampleCases contains the examples that every registered solution must pass.
// Add one entry per example. Encode arguments as a positional JSON array.
var sampleCases = []testrunner.Case{
	{
		Name:  "Example 1",
		Input: `[[2,7,11,15],9]`,
		Want:  `[0,1]`,
	},
	{
		Name:  "Example 2",
		Input: `[[3,2,4],6]`,
		Want:  `[1,2]`,
	},
	{
		Name:  "Example 3",
		Input: `[[3,3],6]`,
		Want:  `[0,1]`,
	},
	{
		Name:  "Minimum length with distinct values",
		Input: `[[1,2],3]`,
		Want:  `[0,1]`,
	},
	{
		Name:  "Zero target with duplicate zeroes",
		Input: `[[0,4,3,0],0]`,
		Want:  `[0,3]`,
	},
	{
		Name:  "Negative target",
		Input: `[[-8,-3,-7,-2],-5]`,
		Want:  `[1,3]`,
	},
	{
		Name:  "Mixed signs and value greater than target",
		Input: `[[10,-2,7,-3],7]`,
		Want:  `[0,3]`,
	},
	{
		Name:  "Answer is the final pair",
		Input: `[[1,2,3,7],10]`,
		Want:  `[2,3]`,
	},
	{
		Name:  "Equal answer values separated by other values",
		Input: `[[1,5,1,9],2]`,
		Want:  `[0,2]`,
	},
	{
		Name:  "Constraint extremes",
		Input: `[[1000000000,-1000000000],0]`,
		Want:  `[0,1]`,
	},
	{
		Name:  "Maximum target boundary",
		Input: `[[1000000000,-1,4],999999999]`,
		Want:  `[0,1]`,
	},
	{
		Name:  "Minimum target boundary",
		Input: `[[-1000000000,0,4],-1000000000]`,
		Want:  `[0,1]`,
	},
	{
		Name:  "Answer uses first and last indices",
		Input: `[[8,1,2,6,-3],5]`,
		Want:  `[0,4]`,
	},
}

// archiveReady is set to true only after the problem, samples, three solutions,
// analysis, and regression coverage are complete and verified.
const archiveReady = true

// solutions lists every implementation that the sample runner checks. Keep the
// current attempt registered while developing it. When Codex adds three final
// approaches, every approach is registered here and receives the same cases.
var solutions = []testrunner.Solution{
	{Name: "current attempt", Run: solveCurrentAttempt},
	{Name: "check every pair", Run: solveByCheckingEveryPair},
	{Name: "sort and use two pointers", Run: solveBySortingAndTwoPointers},
	{Name: "remember complements", Run: solveByRememberingComplements},
}

// solveCurrentAttempt is the only function you need to change initially. It
// receives the exact sample input string and returns the expected output form.
// Use testrunner.DecodeJSONArgs to decode arguments into local variables.
// Leading and trailing whitespace is ignored during comparison; internal
// whitespace remains significant.
func solveCurrentAttempt(input string) string {
	// Parse input, implement your algorithm, and return the formatted answer.
	var nums []int
	var target int
	var ret string

	if err := testrunner.DecodeJSONArgs(input, &nums, &target); err != nil {
		panic(err)
	}

	for i := 0; i < len(nums)-1; i++ {
		curNums := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if target-curNums == nums[j] {
				ret = fmt.Sprintf("[%d,%d]", i, j)
				return ret
			}
		}
	}

	return ret
}

// solveByCheckingEveryPair is the direct baseline. It improves the control flow
// of the current attempt by returning as soon as the unique answer is found.
func solveByCheckingEveryPair(input string) string {
	// Decode the JSON arguments once before running the typed search.
	nums, target := decodeTwoSumInput(input)

	// Choose each index except the last as the first member of a pair.
	for first := 0; first < len(nums)-1; first++ {
		// Start after first so an element is never paired with itself.
		for second := first + 1; second < len(nums); second++ {
			// Check whether the current pair produces the requested target.
			if nums[first]+nums[second] == target {
				// Return immediately because the problem guarantees one answer.
				return formatPair(first, second)
			}
		}
	}

	// Make a violated problem contract visible instead of returning bad output.
	panic("two sum: input has no solution")
}

// indexedNumber keeps a value connected to its position before sorting.
type indexedNumber struct {
	value int // value participates in the sorted two-pointer comparison.
	index int // index is the value's original position in nums.
}

// solveBySortingAndTwoPointers trades linear auxiliary space for an O(n log n)
// ordering step, then eliminates impossible pairs from both ends.
func solveBySortingAndTwoPointers(input string) string {
	// Decode the input without changing the order of the original slice.
	nums, target := decodeTwoSumInput(input)
	// Allocate one sortable record for every original value.
	ordered := make([]indexedNumber, len(nums))
	// Copy both each number and its original index into the sortable slice.
	for index, value := range nums {
		// Retaining index lets the solution return positions after sorting values.
		ordered[index] = indexedNumber{value: value, index: index}
	}

	// Sort by value so pointer movements have a monotonic meaning.
	sort.Slice(ordered, func(first, second int) bool {
		// Break equal-value ties by original index for deterministic ordering.
		if ordered[first].value == ordered[second].value {
			// Place the lower original index first when the values are equal.
			return ordered[first].index < ordered[second].index
		}
		// Otherwise place the smaller numeric value first.
		return ordered[first].value < ordered[second].value
	})

	// Begin with the smallest and largest remaining values.
	for left, right := 0, len(ordered)-1; left < right; {
		// Compute the sum represented by the current pointer positions.
		sum := ordered[left].value + ordered[right].value
		// Compare the sum with the target to decide which side to discard.
		switch {
		// A small sum means the current smallest value cannot be in the answer.
		case sum < target:
			// Move to the next larger candidate on the left.
			left++
		// A large sum means the current largest value cannot be in the answer.
		case sum > target:
			// Move to the next smaller candidate on the right.
			right--
		// Equality means the two records contain the required values.
		default:
			// Return their original indices in ascending index order.
			return formatPair(ordered[left].index, ordered[right].index)
		}
	}

	// Make a violated problem contract visible instead of returning bad output.
	panic("two sum: input has no solution")
}

// solveByRememberingComplements performs one left-to-right pass. The map holds
// only earlier values, so a match always uses two different indices.
func solveByRememberingComplements(input string) string {
	// Decode the JSON arguments once before running the typed search.
	nums, target := decodeTwoSumInput(input)
	// Map every previously visited value to one of its earlier indices.
	indexByValue := make(map[int]int)

	// Visit each value once from left to right.
	for index, value := range nums {
		// Calculate the value needed to complete the target.
		complement := target - value
		// Look up the complement and whether it exists as two explicit results.
		earlierIndex, found := indexByValue[complement]
		// Continue into this branch only when an earlier complement was found.
		if found {
			// The mapped index is earlier, so the two indices are distinct.
			return formatPair(earlierIndex, index)
		}
		// Store only after lookup so a value cannot match itself.
		indexByValue[value] = index
	}

	// Make a violated problem contract visible instead of returning bad output.
	panic("two sum: input has no solution")
}

// decodeTwoSumInput converts the runner's positional JSON into typed values.
func decodeTwoSumInput(input string) ([]int, int) {
	// nums receives the first positional argument.
	var nums []int
	// target receives the second positional argument.
	var target int
	// Decode both arguments and retain context if the adapter input is invalid.
	if err := testrunner.DecodeJSONArgs(input, &nums, &target); err != nil {
		// Invalid runner input is a harness error, so stop with its exact cause.
		panic(err)
	}
	// Return the typed values used by every provided algorithm.
	return nums, target
}

// formatPair produces the deterministic JSON result expected by the runner.
func formatPair(first, second int) string {
	// Normalize index order because the problem accepts either order.
	if first > second {
		// Swap the indices so every solution produces the same output string.
		first, second = second, first
	}
	// Encode the two indices without internal whitespace.
	return fmt.Sprintf("[%d,%d]", first, second)
}

/*
===============================================================================
AGENT REVIEW AND COMPLEXITY ANALYSIS

Current attempt:
- What is correct: The outer loop chooses the first index and the inner loop
  checks only later indices. That prevents using one element twice and checks
  every unordered pair at most once. The subtraction comparison is equivalent
  to nums[i] + nums[j] == target within the stated integer bounds. Returning at
  the first match is valid because the statement guarantees one answer.
- Invariant: Before each inner-loop iteration, every pair ordered before (i, j)
  has been checked and no such pair is an answer. Because j is always greater
  than i, any reported pair contains distinct indices.
- Example walkthrough for nums = [2,7,11,15], target = 9:
  1. Decode nums and target, then start the outer loop with i = 0.
  2. Set curNums = nums[0] = 2 and start the inner loop with j = 1.
  3. Compute target-curNums = 9-2 = 7 and compare it with nums[1] = 7.
  4. The values match, so format and return indices [0,1].
- First bottleneck: In the worst case the answer is near the end, so almost all
  n(n-1)/2 pairs are checked. For n = 10^4 that is about 50 million checks and
  does not satisfy the requested subquadratic follow-up.
- Time complexity: Best O(1) when indices (0, 1) are the answer; average and
  worst O(n^2). The early return improves completed work but not the worst-case
  growth rate.
- Auxiliary-space complexity: O(1) for the search. End to end, JSON decoding
  allocates the O(n) nums slice and fmt.Sprintf allocates the constant-size
  result string.
- Ways to improve: Keep the early return, return the formatted pair directly
  instead of assigning the temporary ret variable, use a singular name such as
  currentNum, separate input/output adaptation from the typed search, and fail
  explicitly instead of returning an empty string if the input contract is
  violated. Use a complement map for expected O(n) time. The added regression
  cases make each of those changes repeatable.
- Relevant concepts: concepts/data-structures/hash-map/README.md for complement
  lookup and concepts/patterns/two-pointers/README.md for sorted pair search.

Solution 1 - Direct pair enumeration (solveByCheckingEveryPair):
- Approach: Enumerate (first, second) with first < second and return immediately
  when their values sum to target. This intentionally formalizes the same
  algorithmic baseline as the current attempt before comparing alternatives.
- Correctness: Every legal pair has exactly one ordering with first < second,
  and the nested loops visit all such orderings. Therefore the unique valid pair
  is visited and returned, while no element can be paired with itself.
- Example walkthrough for nums = [1,2,3,7], target = 10:
  1. With first = 0, test sums 1+2 = 3, 1+3 = 4, and 1+7 = 8.
  2. With first = 1, test sums 2+3 = 5 and 2+7 = 9.
  3. With first = 2 and second = 3, calculate 3+7 = 10.
  4. The sum matches target, so return indices [2,3].
- Time complexity: Best O(1); average and worst O(n^2). The worst case performs
  n(n-1)/2 comparisons.
- Auxiliary-space complexity: O(1) for the algorithm. The shared adapter still
  uses O(n) space to decode the JSON input and O(1) output space.

Solution 2 - Sort indexed values and use two pointers
(solveBySortingAndTwoPointers):
- Approach: Copy each value with its original index, sort the copy by value,
  and move pointers inward according to whether their sum is too small or too
  large. Original indices are retained because sorting changes positions.
- Invariant: At the start of each iteration, any remaining answer lies within
  [left, right]. If the sum is too small, pairing the left value with any value
  no larger than the right value is also too small, so left can be discarded.
  The symmetric argument permits discarding right when the sum is too large.
- Correctness: Each movement removes only a value that cannot participate in a
  remaining answer. The guaranteed answer therefore remains in the interval
  until its two values become the endpoints and are returned.
- Example walkthrough for nums = [2,7,11,15], target = 9:
  1. Build and sort records as [(2,0),(7,1),(11,2),(15,3)], where each pair is
     (value, original index).
  2. Set left at 2 and right at 15. Their sum is 17, which is too large, so move
     right from 15 to 11.
  3. The new sum is 2+11 = 13, which is still too large, so move right to 7.
  4. The sum is now 2+7 = 9, so return original indices [0,1].
- Time complexity: Copying is O(n). Go's adaptive sort can take O(n) on already
  ordered data and takes O(n log n) on average and in the worst case; the scan
  is best O(1) and worst O(n). The complete approach is therefore best O(n) and
  average/worst O(n log n).
- Auxiliary-space complexity: O(n) for the indexed copy, plus implementation
  stack space used by sorting. JSON decoding separately materializes O(n) input.

Solution 3 - One-pass complement map (solveByRememberingComplements):
- Approach: Scan left to right. Before storing nums[index], ask whether
  target-nums[index] was seen earlier. Store values only after lookup so equal
  values such as [3,3] use two distinct indices.
- Invariant: Before processing index i, indexByValue contains an index for every
  distinct value in nums[0:i] and contains no index at or after i.
- Correctness: When the later member of the unique answer is processed, the
  earlier member is already mapped under exactly the needed complement, so the
  algorithm returns their distinct indices. A returned map match always sums to
  target by construction.
- Example walkthrough for nums = [2,7,11,15], target = 9:
  1. Start with an empty indexByValue map.
  2. At index 0, value is 2 and complement is 9-2 = 7. The map has no 7, so
     store 2 -> 0.
  3. At index 1, value is 7 and complement is 9-7 = 2. The map contains 2 -> 0,
     so earlierIndex is 0 and found is true.
  4. Return earlier index 0 and current index 1 as [0,1].
- Time complexity: Best expected O(1) when the first two values match; average
  and worst expected O(n), assuming expected O(1) Go map operations. A
  pathological hash-collision model has theoretical O(n^2) worst-case time.
- Auxiliary-space complexity: Best O(1) when a match is found immediately and
  worst O(n) map entries, in addition to the O(n) decoded input and
  constant-size output.

AGENT REVIEW AND COMPLEXITY ANALYSIS END
===============================================================================
*/

func main() {
	if err := testrunner.Run(os.Stdout, sampleCases, solutions); err != nil {
		fmt.Fprintln(os.Stderr, "workbench:", err)
		os.Exit(1)
	}
}
