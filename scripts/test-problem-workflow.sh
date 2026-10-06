#!/usr/bin/env bash

set -euo pipefail

script_directory=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
repository_root=$(cd "$script_directory/.." && pwd)
temporary_root=$(mktemp -d)
trap 'rm -rf "$temporary_root"' EXIT

mkdir -p "$temporary_root/workbench" "$temporary_root/problems"
cp "$repository_root/go.mod" "$temporary_root/go.mod"
mkdir -p "$temporary_root/internal"
cp -R "$repository_root/internal/testrunner" "$temporary_root/internal/testrunner"
cp "$repository_root/workbench/main.go.template" "$temporary_root/workbench/main.go.template"
cp "$repository_root/workbench/main.go.template" "$temporary_root/workbench/main.go"
cp "$repository_root/problems/README.md" "$temporary_root/problems/README.md"

ALGOS_REPO_ROOT=$temporary_root "$script_directory/reset-workbench.sh"

sed 's/Include the source URL if available./Unarchived test problem./' \
  "$repository_root/workbench/main.go.template" > "$temporary_root/workbench/main.go"
if ALGOS_REPO_ROOT=$temporary_root "$script_directory/reset-workbench.sh" >/dev/null 2>&1; then
  echo "error: unfinished workbench reset unexpectedly succeeded" >&2
  exit 1
fi
ALGOS_REPO_ROOT=$temporary_root "$script_directory/reset-workbench.sh" --force
cmp -s "$temporary_root/workbench/main.go.template" "$temporary_root/workbench/main.go"

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" easy two-sum >/dev/null 2>&1; then
  echo "error: incomplete workbench archive unexpectedly succeeded" >&2
  exit 1
fi

sed 's/^const archiveReady = false$/const archiveReady = true/' \
  "$repository_root/workbench/main.go" > "$temporary_root/workbench/main.go"

ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" easy two-sum

archived_file="$temporary_root/problems/easy_two_sum.go"
test -f "$archived_file"
test "$(sed -n '1p' "$archived_file")" = "//go:build ignore"
rg -q '^// Difficulty: Easy$' "$archived_file"
rg -q 'easy_two_sum.go' "$temporary_root/problems/README.md"

ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh" easy_two_sum.go
ALGOS_REPO_ROOT=$temporary_root "$script_directory/run-problem.sh" problems/easy_two_sum.go
ALGOS_REPO_ROOT=$temporary_root "$script_directory/test-problems.sh"

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" easy two-sum >/dev/null 2>&1; then
  echo "error: duplicate archive unexpectedly succeeded" >&2
  exit 1
fi

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" expert invalid >/dev/null 2>&1; then
  echo "error: invalid difficulty unexpectedly succeeded" >&2
  exit 1
fi

if ALGOS_REPO_ROOT=$temporary_root "$script_directory/archive-workbench.sh" hard invalid--slug >/dev/null 2>&1; then
  echo "error: invalid slug unexpectedly succeeded" >&2
  exit 1
fi

ALGOS_REPO_ROOT=$temporary_root "$script_directory/reset-workbench.sh"
cmp -s "$temporary_root/workbench/main.go.template" "$temporary_root/workbench/main.go"

echo "problem workflow tests passed"
