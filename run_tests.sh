#!/bin/bash
set -e

timestamp=$(date +"%Y-%m-%dT%H-%M-%S")
report_dir="test_reports/$timestamp"
mkdir -p "$report_dir"

echo "Running all tests with coverage..."
go test -v -coverprofile="$report_dir/coverage.out" -json ./... > "$report_dir/report.json"

echo "Generating HTML coverage report..."
gocov convert "$report_dir/coverage.out" | gocov-html > "$report_dir/coverage.html" 2>/dev/null || go tool cover -html="$report_dir/coverage.out" -o "$report_dir/coverage.html"

echo "Test run complete. Reports stored in $report_dir"
