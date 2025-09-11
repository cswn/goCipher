init:
	$(GO) version
	$(GO) mod download
	$(GO) mod tidy

build:
	go build -o ./build/goCipher ./cmd

test/unit:
	go test ./cmd