.PHONY : format install build

run-this:
	echo "hello"

everything-oke:
	go run ./bin/app/main.go

install:
	go mod download

build:
	docker-compose up -d

down:
	docker-compose down

start:
	./main
