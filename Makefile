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

image:
	docker build -t romi666/komikku-api:1.0 .

push:
	docker push romi666/komikku-api:1.0

start:
	./main
