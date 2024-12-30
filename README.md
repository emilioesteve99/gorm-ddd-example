# Project Name

## Description
This is a simple API that allows to create, read, update, and delete users

It is built using Go and PostgresSQL

It's been implemented using SOLID principles and Domain Driven Design

## Prerequisites
- **Go 1.23.4**
- **Docker**
- **Docker Compose**

## Setup

1 **Install Go dependencies:**
    ```go mod download```

2 **Run Docker Compose:**
    ```docker-compose up --build```

## Usage
There is a postman collection in the `.postman` directory that can be used to test the API

## Running Tests
To run tests, use the following command: ```go test ./...```