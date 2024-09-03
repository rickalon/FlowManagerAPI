run: build
	./bin/fmapi

build: 
	go build -o ./bin/fmapi cmd/FlowManagerAPI/main.go

test:
	go test -v ./tests/...