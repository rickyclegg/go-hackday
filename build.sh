#!/usr/bin/env bash

set -e

echo "Running build for go-hackday"
go test *.go
echo "All tests pass"
docker build -t=go-hackday .
echo "Complete"
