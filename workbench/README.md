# Active Problem Workbench

This directory holds one active problem at a time. You only edit
[`main.go`](main.go). It contains labeled sections for the statement, sample
cases, current attempt, final solutions, and complexity analysis. Universal
sample execution lives in [`internal/testrunner`](../internal/testrunner/).

## Quick start

1. Paste the problem into the `PROBLEM` block comment in `main.go`.
2. Define any judge-owned types and problem-specific input/output adapters in
   the section immediately below the statement.
3. Add every supplied example to the `sampleCases` slice.
4. Implement `solveCurrentAttempt` using the typed adapter values.
5. Run `./scripts/workbench.sh` from the repository root.
6. Ask Codex to review or solve "the problem in the workbench."

When you ask Codex to add the inputs or examples, it also completes the initial
input setup in your current implementation: it creates or updates the
problem-specific parser and calls it at the start of `solveCurrentAttempt`, so
your algorithm begins with strongly typed arguments. If structural output needs
conversion, Codex also wires the matching formatter without replacing your
algorithm body.

The agent sets `archiveReady` to `true` only after the completed problem and all
three solutions pass their checks. The archive script refuses incomplete
workbench content.

The universal workbench function accepts raw sample input and returns raw
output. For a function-style judge such as LeetCode, encode the arguments in
call order as a JSON array:

```go
Input: `[[2,7,11,15],9]`,
Want:  `[0,1]`,
```

For simple values, a problem-specific parser can decode each argument directly
and return strongly typed values:

```go
func parseTwoSumInput(input string) ([]int, int) {
	var nums []int
	var target int
	if err := testrunner.DecodeJSONArgs(input, &nums, &target); err != nil {
		panic(fmt.Errorf("parse two sum input: %w", err))
	}
	return nums, target
}
```

Structural problems keep the same readable sample representation and perform
the conversion inside their adapter. For example, a linked-list parser decodes
two `[]int` arguments and builds two `*ListNode` values. Trees, graphs, matrices,
and custom judge types follow the same pattern. Add a problem-specific formatter
when the typed result also needs conversion back to the statement's output
form. Handle decoding and formatting errors at these adapter boundaries.

If copied problem text contains the uncommon sequence `*/`, change it to `* /`
so it does not close the Go block comment early.

Codex will preserve and review the attempt, connect it to relevant material in
[`concepts/`](../concepts/), register three distinct solutions, and apply the
same sample and regression cases to all three. Detailed correctness and
complexity notes stay in the analysis block in `main.go`.

Once verified, the complete problem is copied to its permanent location under
[`problems/`](../problems/) with
`./scripts/archive-workbench.sh <easy|medium|hard> <descriptive-slug>`. Do not
overwrite an unarchived workbench problem.

After verifying the archived file, restore the canonical blank template with:

```bash
./scripts/reset-workbench.sh
```

The reset command refuses unfinished changes. Use `--force` only when you
explicitly want to discard the current workbench without archiving it.
