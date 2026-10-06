#!/usr/bin/env bash

set -euo pipefail

usage() {
  echo "usage: $0 <difficulty_problem_name.go>" >&2
}

fail() {
  echo "error: $1" >&2
  exit 1
}

if [[ $# -ne 1 ]]; then
  usage
  exit 2
fi

requested_file=$1
case $requested_file in
  problems/*) filename=${requested_file#problems/} ;;
  */*) fail "provide a filename or problems/<filename>" ;;
  *) filename=$requested_file ;;
esac

[[ $filename =~ ^(easy|medium|hard)_[a-z0-9]+(_[a-z0-9]+)*\.go$ ]] || fail "filename must match <easy|medium|hard>_<descriptive_slug>.go"

script_directory=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
default_root=$(cd "$script_directory/.." && pwd)
repository_root=${ALGOS_REPO_ROOT:-$default_root}
problem_file="$repository_root/problems/$filename"

[[ -f $problem_file ]] || fail "$problem_file was not found"

(cd "$repository_root" && go run "$problem_file")
