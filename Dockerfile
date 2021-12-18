FROM golang:1.16-alpine

LABEL maintainer="Romi666"

WORKDIR /usr/src/app

COPY . .

RUN go mod download \
  && go build -tags musl -o main ./bin/app

# Expose port
EXPOSE 3011

# Run application
CMD ["./main"]
