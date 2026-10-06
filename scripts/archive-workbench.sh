#!/usr/bin/env bash

set -euo pipefail

usage() {
  echo "usage: $0 <easy|medium|hard> <descriptive-slug>" >&2
}

fail() {
  echo "error: $1" >&2
  exit 1
}

if [[ $# -ne 2 ]]; then
  usage
  exit 2
fi

difficulty=$1
slug=$2

case $difficulty in
  easy) difficulty_label=Easy ;;
  medium) difficulty_label=Medium ;;
  hard) difficulty_label=Hard ;;
  *) fail "difficulty must be easy, medium, or hard" ;;
esac

[[ $slug =~ ^[a-z0-9]+(-[a-z0-9]+)*$ ]] || fail "slug must contain lowercase letters, digits, and single hyphen separators"

script_directory=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
default_root=$(cd "$script_directory/.." && pwd)
repository_root=${ALGOS_REPO_ROOT:-$default_root}
source_file="$repository_root/workbench/main.go"
problem_root="$repository_root/problems"
filename="${difficulty}_${slug//-/_}.go"
destination="$problem_root/$filename"

[[ -f $source_file ]] || fail "$source_file was not found"
[[ -f "$problem_root/README.md" ]] || fail "$problem_root/README.md was not found"
[[ ! -e $destination ]] || fail "$destination already exists"
rg -q '^const archiveReady = true$' "$source_file" || fail "workbench is not marked ready to archive"

temporary_directory=$(mktemp -d "$repository_root/.archive-workbench.XXXXXX")
temporary_file="$temporary_directory/$filename"
cleanup() {
  rm -f "$temporary_file"
  rmdir "$temporary_directory" 2>/dev/null || true
}
trap cleanup EXIT

{
  printf '%s\n\n' '//go:build ignore'
  printf '// Difficulty: %s\n\n' "$difficulty_label"
  sed -n '1,$p' "$source_file"
} > "$temporary_file"

gofmt -w "$temporary_file"
(cd "$repository_root" && go run "$temporary_file")

mv "$temporary_file" "$destination"
printf -- '- [%s](%s)\n' "$filename" "$filename" >> "$problem_root/README.md"
rmdir "$temporary_directory"
trap - EXIT

echo "archived $destination"
echo "after verification, reset with ./scripts/reset-workbench.sh"
