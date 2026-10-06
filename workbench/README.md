# Active Problem Workbench

This directory holds one active problem at a time. You only edit
[`main.go`](main.go). It contains labeled sections for the statement, sample
cases, current attempt, final solutions, and complexity analysis. Universal
sample execution lives in [`internal/testrunner`](../internal/testrunner/).

## Quick start

1. Paste the problem into the `PROBLEM` block comment in `main.go`.
2. Add every supplied example to the `sampleCases` slice.
3. Implement `solveCurrentAttempt`.
4. Run `./scripts/workbench.sh` from the repository root.
5. Ask Codex to review or solve "the problem in the workbench."

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

Decode each argument directly into a local variable, then keep the algorithm
itself normally typed:

```go
var nums []int
var target int
err := testrunner.DecodeJSONArgs(input, &nums, &target)
```

Handle the returned error at the adapter boundary. If copied problem text
contains the uncommon sequence `*/`, change it to `* /` so it does not close
the Go block comment early.

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
