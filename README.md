# Monorepo

> A monorepo integrating Kratos, Ent, and GraphQL for building modern, scalable, and maintainable backend services.

## 🚀 Overview

keg (short for Kratos, Ent, and GraphQL) is a highly modular monorepo structure that provides a robust foundation for developing backend services in Go. By combining the best practices of Kratos (a Go microservices framework), Ent (an ORM for Go), and GraphQL, this monorepo streamlines the process of building and scaling backend applications.

## 🎯 Features
* Kratos Integration: Use Kratos for building efficient and scalable microservices with advanced features like service discovery, logging, and configuration management.
* Ent ORM: A powerful, type-safe ORM that simplifies database schema management and querying.
* GraphQL Support: First-class GraphQL support for defining, querying, and mutating data.
* Monorepo Structure: Centralized codebase for better dependency management and shared utilities across services.

## 📂 Project Structure
```
.
├── Makefile
├── bin
├── cmd
│   ├── admin # admin service (example)
│   │   ├── main.go
│   │   ├── wire.go
│   │   └── wire_gen.go
│   └── gateway # kratos-gateway
│       └── main.go
├── configs
│   ├── admin.yaml
│   └── gateway.yaml
├── deploy
│   ├── components
│   │   └── docker-compose.yaml
│   ├── build
│   │   ├── Dockerfile
│   │   ├── admin
│   │   │   └── Dockerfile # Dockerfile for admin service
│   │   └── common
│   │       └── Dockerfile # common Dockerfile for all services
│   └── kubernetes # k8s deployment template files
│       ├── admin.yaml
│       ├── configmap.yaml
│       ├── gateway.yaml
│       └── output # k8s deployment files output directory
├── internal
│   ├── app
│   │   ├── admin
│   │   │   ├── biz
│   │   │   ├── cmd
│   │   │   ├── conf
│   │   │   ├── data
│   │   │   ├── job
│   │   │   ├── schedule
│   │   │   ├── server
│   │   │   └── service
│   │   │       ├── graphql
│   │   │       │   ├── dataloader # Optimizing N+1 database queries
│   │   │       │   │   └── dataloader.go
│   │   │       │   ├── ent.graphql
│   │   │       │   ├── ent.resolvers.go
│   │   │       │   ├── generate.go
│   │   │       │   ├── gqlgen.yml
│   │   │       │   ├── model
│   │   │       │   │   └── models_gen.go
│   │   │       │   ├── resolver.go
│   │   │       └── service.go
│   │   ├── gateway
│   │   │   └── middleware # custom gateway middleware
│   │   │       ├── auth.go
│   │   │       └── permission.go
│   ├── data # database schema
│   │   ├── example
│   │   │   └── ent
│   │   │       └── schema
│   └── pkg
│       └── auth
│           └── auth.go
├── keg.yaml # keg 配置文件
├── logs
├── scripts
│   ├── base.mk
│   ├── compose.mk
│   └── init.mk
└── website
```