DEFAULT_GOAL := build

build:
	go build -ldflags "-s -w" -o bin/prisma.exe ./cmd/api/main.go

#load-test:
#	artillery run --output ./internal/tests/load-test/result.json ./internal/tests/load-test.yaml