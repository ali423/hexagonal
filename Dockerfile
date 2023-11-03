FROM golang:1.21.1-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o shotener cmd/shotener/main.go

CMD ["./shotener"]
