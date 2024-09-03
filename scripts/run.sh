#!/bin/bash
echo "Deleting previous binaries"
rm -r ./bin/
echo "Building API"
go build -o ./bin/FlowManagerAPI/flapi ./cmd/FlowManagerAPI/
echo "Exec API"
./bin/FlowManagerAPI/flapi