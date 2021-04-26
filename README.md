# Prerequisites

- Install the protocol buffer compiler `protoc`

# Overview

- The project will implement the following pattern, repository, factory, model, domain and services.

# Questions

- Not sure exactly sure what to put into the repository struct, should the signature be the same as the grpc service definition? Then that would be a unnecessary duplication, no?

# TODO

- Add UserService, reimplement the ArticleService so they interact with each other together (see how it is implemented in the WildWorkout project)
- Deploy the services and write terraform code for it
- Stage the client and the server
- Create a PWA (NextJS)

# Upcoming Features

- Add http proxy (grpc plugin)
- Add TLS to server
- Server authentication
- Use an asynchronous queue to store multiple articles, batch the article requests and send the batch requests to the server.

Monitoring

- Add open tracing
- Add `zapper` for logging
