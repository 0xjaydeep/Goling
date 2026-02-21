#!/usr/bin/env bash
# verify.sh — Check your Golings exercise progress
#
# Usage:
#   bash verify.sh          # check all exercises
#   bash verify.sh 01       # check a specific exercise by number
#
# Exit codes:
#   0 — all exercises pass (or specific exercise passes)
#   1 — one or more exercises fail or are unfinished

set -euo pipefail

# ── colours ──────────────────────────────────────────────────────────────────
if [[ -t 1 ]]; then
  GREEN="\033[0;32m"
  YELLOW="\033[0;33m"
  RED="\033[0;31m"
  BOLD="\033[1m"
  RESET="\033[0m"
else
  GREEN="" YELLOW="" RED="" BOLD="" RESET=""
fi

# ── helpers ───────────────────────────────────────────────────────────────────
pass()  { printf "  ${GREEN}✓ PASS${RESET}   %s\n" "$1"; }
todo()  { printf "  ${YELLOW}○ TODO${RESET}   %s\n" "$1"; }
fail()  { printf "  ${RED}✗ FAIL${RESET}   %s\n" "$1"; }

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

EXERCISES=(
  "exercises/01-calculator"
  "exercises/02-phonebook"
  "exercises/03-word-frequency"
  "exercises/04-student-grades"
  "exercises/05-safe-stack"
  "exercises/06-mini-bank"
)

# ── filter by argument (optional) ────────────────────────────────────────────
FILTER="${1:-}"
if [[ -n "$FILTER" ]]; then
  FILTERED=()
  for ex in "${EXERCISES[@]}"; do
    if [[ "$ex" == *"$FILTER"* ]]; then
      FILTERED+=("$ex")
    fi
  done
  EXERCISES=("${FILTERED[@]}")
  if [[ ${#EXERCISES[@]} -eq 0 ]]; then
    echo "No exercise matching \"$FILTER\" found."
    exit 1
  fi
fi

# ── header ────────────────────────────────────────────────────────────────────
printf "\n${BOLD}╔══════════════════════════════════════╗${RESET}\n"
printf "${BOLD}║        Golings Progress Check        ║${RESET}\n"
printf "${BOLD}╚══════════════════════════════════════╝${RESET}\n\n"

PASS=0
TODO_COUNT=0
FAIL=0

# ── check each exercise ───────────────────────────────────────────────────────
for ex in "${EXERCISES[@]}"; do
  name=$(basename "$ex")
  dir="$SCRIPT_DIR/$ex"

  if [[ ! -f "$dir/main.go" ]]; then
    fail "$name  (main.go not found)"
    ((FAIL++)) || true
    continue
  fi

  # Still has unimplemented stubs?
  if grep -q "not implemented" "$dir/main.go" 2>/dev/null; then
    todo "$name"
    ((TODO_COUNT++)) || true
    continue
  fi

  # Try to compile and run
  output=$(cd "$dir" && go run main.go 2>&1) && run_ok=true || run_ok=false

  if $run_ok; then
    pass "$name"
    ((PASS++)) || true
  else
    fail "$name"
    # Print first 3 lines of error output, indented
    echo "$output" | head -3 | sed 's/^/    /'
    ((FAIL++)) || true
  fi
done

# ── summary ───────────────────────────────────────────────────────────────────
printf "\n${BOLD}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${RESET}\n"
printf "  ${GREEN}Done: $PASS${RESET}   ${YELLOW}Todo: $TODO_COUNT${RESET}   ${RED}Fail: $FAIL${RESET}\n"
printf "${BOLD}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${RESET}\n"

# ── celebrate completion ──────────────────────────────────────────────────────
if [[ $TODO_COUNT -eq 0 && $FAIL -eq 0 && $PASS -gt 0 ]]; then
  printf "\n${GREEN}${BOLD}  All exercises complete!${RESET}\n"
  printf "  You have finished Golings. You know Go basics.\n\n"
  printf "  Next steps:\n"
  printf "    → Goroutines & Channels (go.dev/tour/concurrency)\n"
  printf "    → go test  (table-driven tests)\n"
  printf "    → net/http (build a web server in ~10 lines)\n\n"
fi

# ── exit code ─────────────────────────────────────────────────────────────────
[[ $FAIL -eq 0 ]]
