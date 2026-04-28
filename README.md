# Go gRPC Microservices Pet Project

A lightweight microservices architecture built with **Go**, demonstrating inter-service communication using **gRPC** and **Protocol Buffers**.

## Architecture
The system consists of two decoupled services:
* **Stock Service**: A gRPC server that manages inventory logic and checks item availability.
* **Order Service**: A gRPC client that simulates order placement by communicating with the Stock Service.

## Tech Stack
* **Language:** Go (Golang)
* **Communication:** gRPC
* **Data Serialization:** Protocol Buffers (proto3)
* **Development Environment:** VS Code

## Getting Started

### 1. Prerequisites
Ensure you have the Protocol Buffer compiler (`protoc`) installed along with the Go plugins:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest