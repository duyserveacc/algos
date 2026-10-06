# Problem Index

This is a flat archive of individually runnable Go programs. Each file contains
its statement, samples, three solutions, and correctness and complexity
analysis. Universal execution and comparison behavior comes from
[`internal/testrunner`](../internal/testrunner/).

Filenames use `<difficulty>_<descriptive_slug>.go`:

- `easy_two_sum.go`
- `medium_coin_change.go`
- `hard_edit_distance.go`

Source platforms do not affect organization. A source URL can remain inside
the file for attribution. Every archived file starts with `//go:build ignore`,
which prevents directory-wide Go commands from combining independent `main`
programs.

Archive and run problems with:

```bash
./scripts/archive-workbench.sh easy two-sum
./scripts/run-problem.sh easy_two_sum.go
./scripts/test-problems.sh
```

## Problems
