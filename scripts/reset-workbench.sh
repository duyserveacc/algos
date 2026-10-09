#!/usr/bin/env bash

set -euo pipefail

usage() {
  echo "usage: $0 [--force]" >&2
}

fail() {
  echo "error: $1" >&2
  exit 1
}

force=false
if [[ $# -gt 1 ]]; then
  usage
  exit 2
fi
if [[ $# -eq 1 ]]; then
  [[ $1 == "--force" ]] || {
    usage
    exit 2
  }
  force=true
fi

script_directory=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
default_root=$(cd "$script_directory/.." && pwd)
repository_root=${ALGOS_REPO_ROOT:-$default_root}
template_file="$repository_root/workbench/main.go.template"
workbench_file="$repository_root/workbench/main.go"

[[ -f $template_file ]] || fail "$template_file was not found"

if [[ -f $workbench_file ]] && cmp -s "$template_file" "$workbench_file"; then
  echo "workbench is already reset"
  exit 0
fi

if [[ -f $workbench_file ]] && ! rg -q '^const archiveReady = true$' "$workbench_file" && [[ $force != true ]]; then
  fail "workbench contains unarchived changes; archive it first or rerun with --force"
fi

temporary_directory=$(mktemp -d "$repository_root/.reset-workbench.XXXXXX")
temporary_file="$temporary_directory/main.go"
temporary_test_files=()
cleanup() {
  if [[ ${#temporary_test_files[@]} -gt 0 ]]; then
    rm -f "${temporary_test_files[@]}"
  fi
  rm -f "$temporary_file"
  rmdir "$temporary_directory" 2>/dev/null || true
}
trap cleanup EXIT

cp "$template_file" "$temporary_file"
gofmt -w "$temporary_file"
(cd "$repository_root" && go run "$temporary_file")

if [[ -f "$repository_root/go.mod" ]]; then
  shopt -s nullglob
  workbench_test_files=("$repository_root"/workbench/*_test.go)
  shopt -u nullglob

  for test_file in "${workbench_test_files[@]}"; do
    temporary_test_file="$temporary_directory/$(basename "$test_file")"
    cp "$test_file" "$temporary_test_file"
    temporary_test_files+=("$temporary_test_file")
  done

  if [[ ${#temporary_test_files[@]} -gt 0 ]]; then
    (cd "$repository_root" && go test "$temporary_file" "${temporary_test_files[@]}")
  fi
fi

if [[ ${#temporary_test_files[@]} -gt 0 ]]; then
  rm -f "${temporary_test_files[@]}"
  temporary_test_files=()
fi
mv "$temporary_file" "$workbench_file"

echo "reset $workbench_file"
