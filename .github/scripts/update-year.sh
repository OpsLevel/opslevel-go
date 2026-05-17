#!/bin/bash
set -e

# Get current year
CURRENT_YEAR=$(date +%Y)

# Extract year from go.mod module path (portable across macOS and Linux)
MODULE_YEAR=$(grep 'module github.com/opslevel/opslevel-go/v' go.mod | sed -E 's/.*\/v([0-9]{4}).*/\1/' || echo "")

if [ -z "$MODULE_YEAR" ]; then
  echo "Error: Could not extract year from go.mod module path"
  exit 1
fi

echo "Current year: $CURRENT_YEAR"
echo "Module year: $MODULE_YEAR"

if [ "$MODULE_YEAR" == "$CURRENT_YEAR" ]; then
  echo "Module year is already current. No updates needed."
  exit 0
fi

echo "Updating module path from v$MODULE_YEAR to v$CURRENT_YEAR..."

# Update go.mod module path
sed -i "s|github.com/opslevel/opslevel-go/v$MODULE_YEAR|github.com/opslevel/opslevel-go/v$CURRENT_YEAR|g" go.mod

# Update all Go files (primarily test imports)
find . -name "*.go" -type f -exec sed -i "s|github.com/opslevel/opslevel-go/v$MODULE_YEAR|github.com/opslevel/opslevel-go/v$CURRENT_YEAR|g" {} +

# Update README.md
if [ -f README.md ]; then
  sed -i "s|opslevel-go/v$MODULE_YEAR|opslevel-go/v$CURRENT_YEAR|g" README.md
fi

# Run go mod tidy to update go.sum
go mod tidy

echo "Year update complete: v$MODULE_YEAR -> v$CURRENT_YEAR"
echo "updated=true" >> $GITHUB_OUTPUT
