# A simple product API using Golang

This application is created using Gin, a simple Golang API framework and PostgreSQL Database.

## Running with Docker

Requirements: Docker and Docker Compose

```bash
$ docker-compose up -d
```

Use helpers/migrations.sql to create table on database

## Running without Docker

Requirements: Golang and PostgreSQL running

Setup credentials on db/conn.go

```bash
$ go mod download
$ go run cmd/main.go
```
