DEFAULT_GOAL := build-and-run

.PHONY: build-and-run
build-and-run:
	make build
	make run

.PHONY: build
build:
	go build -ldflags "-s -w" -o ./bin/prisma.exe ./cmd/api/main.go

.PHONY: run
run:
	./bin/prisma.exe



#load-test:
#	artillery run --output ./internal/tests/load-test/result.json ./internal/tests/load-test.yaml