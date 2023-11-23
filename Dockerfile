FROM golang:1.21.1-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

# Copy the .env.docker file
COPY .env.docker .env

RUN go build -o shotener cmd/shotener/main.go

CMD ["sh", "-c", "sleep 3 && ./shotener"]