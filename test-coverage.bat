@echo off
set COVERAGE_FILE="coverage.out"
go test -coverprofile=%COVERAGE_FILE% ./...
go tool cover -html=%COVERAGE_FILE%