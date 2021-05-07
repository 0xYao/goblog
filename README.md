# Prerequisites

- Install the [protocol buffer](https://grpc.io/docs/protoc-installation/) compiler `protoc`
- Install [air](https://github.com/cosmtrek/air) for live reload

# Get Started

Assuming the `air` package is installed, to start the app, e.g. all internal services, run `air` in the root project directory. The `clients` folder contains the code that interacts with server, for example, try `go run clients/user_client/main.go` to interact with the user grpc server.

# Overview

- The project will implement the following pattern, repository, factory, model, domain and services.

# Questions

- Not sure exactly sure what to put into the repository struct, should the signature be the same as the grpc service definition? Then that would be a unnecessary duplication, no?

# TODO

- Debug the segmentation error when calling the `article_client/main.go`
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
