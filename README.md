# Prerequisites

- Install the protocol buffer compiler `protoc`

# Get Started

All services are inside the `internal` directory. For example, to start the `ArticleService` from the root directory, run `go run internal/articles/main.go`. Then initialize the client to invocate service methods with `go run clients/article_client/main.go`

# Overview

- The project will implement the following pattern, repository, factory, model, domain and services.

# Questions

- Not sure exactly sure what to put into the repository struct, should the signature be the same as the grpc service definition? Then that would be a unnecessary duplication, no?

# TODO

- Reimplement the ArticleService so it interacts with the UserService (see how it is implemented in the WildWorkout project)
- Implement the reposeitory using MongoDB
- Deploy the services and the db and write terraform code for it
- Stage the client and the server
  - Environment variables management
- Create a PWA (NextJS)

# Upcoming Features

- Add http proxy (grpc plugin)
- Add TLS to server
- Server authentication

Monitoring

- Add open tracing
- Add `zapper` for logging
