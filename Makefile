.PHONY : format install build

run-this:
	echo "hello"

everything-oke:
	go run ./bin/app/main.go

install:
	go mod download

build:
	go build -tags musl -o main ./bin/app

start:
	./main
