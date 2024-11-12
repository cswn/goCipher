init:
	$(GO) version
	$(GO) mod download
	$(GO) mod tidy

build:
	go build -o ./build/main ./cmd

test/unit:
	go test ./cmd