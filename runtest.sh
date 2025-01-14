#!/bin/bash

# Print working directory before running tests in each module
print_and_test() {
    local dir=$1
    echo "Running tests in: $dir"
    cd "$dir" || exit
    go test ./... -v
    cd - > /dev/null || exit
}

# Find all directories containing go.mod
find . -name "go.mod" -exec dirname {} \; | while read -r dir; do
    print_and_test "$dir"
done

# Print summary
echo "===================="
echo "Tests completed in all modules"
