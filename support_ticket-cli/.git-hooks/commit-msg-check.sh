#!/bin/bash
set -e

msg_file="$1"
commit_msg=$(cat "$msg_file")

# Simple Conventional Commits check
if ! echo "$commit_msg" | grep -Eq '^(feat|fix|docs|style|refactor|perf|test|chore)(\(.+\))?: .+'; then
  echo "❌ Commit message is NOT following Conventional Commits:"
  echo "Examples:"
  echo "  feat(login): add login endpoint"
  echo "  fix(api): handle nil pointer"
  echo "  docs(readme): update usage"
  exit 1
fi

echo "✓ Commit message validated successfully"
