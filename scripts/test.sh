#!/bin/bash
echo "Testing API"
go test -v ../tests/
echo "Coverage API"
go test -coverprofile ../tests/coverage.out
got tool cover -html=../tests/coverage.out

