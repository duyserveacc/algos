# Repository Agent Instructions

## Purpose

This repository teaches algorithms and data structures in Go. Keep educational
concept implementations under `concepts/`, the active exercise under
`workbench/`, and completed judge-ready solutions under `problems/`.

## Active Problem Workflow

When the user asks for help with, an answer to, or a review of "the problem in
the workbench," treat `workbench/main.go` as the single source of truth for the
statement, samples, current attempt, solutions, and analysis.

1. Read `workbench/main.go`, `workbench/main_test.go`, and the relevant concept
   pages before changing code.
2. Keep judge-owned types and problem-specific parsing and formatting adapters
   in the labeled section before `sampleCases`. Preserve readable statement
   serialization in the cases, convert structural values inside strongly typed
   problem-named adapters, and keep `internal/testrunner` problem-agnostic.
3. Preserve and review the user's attempt. Explain what is correct, the first
   failing assumption or bottleneck, and how the relevant invariant guides the
   correction.
4. Link the problem to existing pages under `concepts/`. Add or improve a
   concept page when the necessary reference does not exist.
5. Produce three distinct correct approaches when solving: a direct or brute
   force approach, an improved approach, and the preferred approach. An
   educational approach may exceed judge limits, but it must still be correct
   for inputs it can finish. Do not manufacture meaningless variants.
6. Give each approach a descriptive function name and register it in the
   `solutions` slice in `workbench/main.go`. Keep the judge-required function as
   a thin wrapper around the preferred approach when appropriate.
7. Run the same `sampleCases` and table-driven regression cases against all
   three approaches.
   Include supplied examples, minimum or empty inputs when valid, boundary
   conditions, duplicate or adversarial values, and at least one case that
   distinguishes a common incorrect solution.
8. In the analysis section of `workbench/main.go`, explain each approach, its
   correctness argument, and its time and auxiliary-space complexity. State
   best, average, and worst cases when they differ, and account for recursion,
   preprocessing, output, and meaningful Go allocations.
9. Run `./scripts/workbench.sh` and `go test ./...` before claiming that the
   solutions pass. Set `archiveReady` to `true` only after these checks pass and
   all three solutions and their analysis are complete.

Do not replace a non-template workbench problem with a new one until its work
has been archived or the user explicitly says to discard it.

## Archiving

After a problem is complete:

1. Assign `easy`, `medium`, or `hard` based on the required insight,
   implementation complexity, edge cases, and constraints.
2. Run `./scripts/archive-workbench.sh <difficulty> <descriptive-slug>`.
3. Confirm the new archived file under `problems/` runs with
   `./scripts/run-problem.sh <filename>`.
4. Run `./scripts/test-problems.sh` and `go test ./...`.
5. Reset the active file with `./scripts/reset-workbench.sh` only after the
   archive and checks succeed.

Use flat filenames such as `easy_two_sum.go`; do not group by platform or
create per-problem directories. Each archived file contains all
problem-specific material: its statement, samples, three solutions, and
analysis. It imports only the shared `internal/testrunner` infrastructure;
reference concept pages in comments rather than importing their code.

Never use `reset-workbench.sh --force` unless the user explicitly authorizes
discarding the unfinished workbench content.

`workbench/main.go.template` is the canonical blank workbench. When changing
its problem-facing API, keep it synchronized with `workbench/main.go` while
preserving any active problem, samples, solutions, and analysis. Universal
runner behavior and tests belong in `internal/testrunner`.
