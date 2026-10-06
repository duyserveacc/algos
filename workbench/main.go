package main

import (
	"fmt"
	"os"

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


Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

PASTE THE COMPLETE PROBLEM ABOVE THIS LINE
===============================================================================
*/

// sampleCases contains the examples that every registered solution must pass.
// Add one entry per example. Encode function arguments as one JSON value.
var sampleCases = []testrunner.Case{
	{
		Name:  "Example 1",
		Input: `{"nums":[2,7,11,15],"target":9}`,
		Want:  `[0,1]`,
	},
	{
		Name:  "Example 2",
		Input: `{"nums":[3,2,4],"target":6}`,
		Want:  `[1,2]`,
	},
	{
		Name:  "Example 3",
		Input: `{"nums":[3,3],"target":6}`,
		Want:  `[0,1]`,
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
// Use testrunner.DecodeJSON to convert JSON input into a problem-specific type.
// Leading and trailing whitespace is ignored during comparison; internal
// whitespace remains significant.
func solveCurrentAttempt(input string) string {
	// Parse input, implement your algorithm, and return the formatted answer.
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
