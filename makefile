.PHONY: repl build test

repl: build
	@./build/repl

build:
	@go build -o ./build/repl .

test:
	@go test ./lexer

