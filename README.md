# Hexagonal: A Hexagonal Architecture URL Shortener

## Overview

Hexagonal is a simple URL shortener project designed to demonstrate the Hexagonal Architecture in a Golang application. The Hexagonal Architecture, also known as Ports and Adapters or Onion Architecture, is an architectural pattern that aims to create loosely coupled application components.

## Features

- URL Shortening: Generate short URLs for long ones.
- Hexagonal Architecture: Demonstrates a modular and scalable architecture.
- Database Flexibility: Connect to multiple SQL databases seamlessly.

## Tests

This project includes tests. To run tests, use the below command : 

```bash
go test -v -cover ./...
```

## Prerequisites

- [Go](https://golang.org/) installed
- [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/install/) (optional)

## Running the Project

### Using Docker:

```bash
docker-compose up
```

### Using Go Command:

```bash
cp .env.example .env

go run cmd/shotener/*.go
```

