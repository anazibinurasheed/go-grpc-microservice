Go-grpc-microservice
===



## Overview

An example of small e-commerce system using Microservice Architecture. Developed 3 Microservices and 1 API Gateway which handles incoming HTTP requests. The HTTP requests will be forwarded to these Microservices using gRPC.
##
## Note: It's a Monorepo contains 1 api-gateway and 3 services.
##

## Application Infrastructure
## 1. API Gateway:  
Handles incoming HTTP requests

## 2. Auth Service: 
Provides features such as Register, Login and generates Token by JWT and a feature called Validate to validate the token.

## 3. Product Service:
 Provides features such as Add Product, Decrease Stock and Find Product.

## 4. Order Service: 
The only feature we ship in this Microservice is Create Order.

## Technology/Architecture
- Postgresql
- gRPC
- Protocol buffers
- JWT
- Microservice Architecture

## Configurations

Each standalone project have its own config package and env file make changes according to you.


## Create Database 

Create database for each service.
  ```
$ psql postgres
$ CREATE DATABASE auth_svc;
$ CREATE DATABASE order_svc;
$ CREATE DATABASE product_svc;
$ \l
$ \q
  ```


## Install Dependencies 
- Install make
  ```
  sudo apt update

  sudo apt install make  
  ```
- Install go deps
  ```
  make tidy
  ```

- for run
  ```
  make server
  ```


## Directory Hierarchy
```
|—— LICENSE
|—— go-grpc-api-gateway
|    |—— .gitignore
|    |—— Makefile
|    |—— cmd
|        |—— main.go
|    |—— go.mod
|    |—— go.sum
|    |—— notes.txt
|    |—— pkg
|        |—— auth
|            |—— client.go
|            |—— middleware.go
|            |—— pb
|                |—— auth.pb.go
|                |—— auth.proto
|                |—— auth_grpc.pb.go
|            |—— routes
|                |—— login.go
|                |—— register.go
|            |—— routes.go
|        |—— config
|            |—— config.go
|            |—— envs
|                |—— dev.env
|        |—— order
|            |—— client.go
|            |—— pb
|                |—— order.pb.go
|                |—— order.proto
|                |—— order_grpc.pb.go
|            |—— routes
|                |—— create_order.go
|            |—— routes.go
|        |—— product
|            |—— client.go
|            |—— pb
|                |—— product.pb.go
|                |—— product.proto
|                |—— product_grpc.pb.go
|            |—— routes
|                |—— create_product.go
|                |—— find_one.go
|            |—— routes.go
|—— go-grpc-auth-svc
|    |—— .gitignore
|    |—— Makefile
|    |—— cmd
|        |—— main.go
|    |—— go.mod
|    |—— go.sum
|    |—— notes.txt
|    |—— pkg
|        |—— config
|            |—— config.go
|            |—— envs
|                |—— dev.env
|        |—— db
|            |—— db.go
|        |—— models
|            |—— auth.go
|        |—— pb
|            |—— auth.pb.go
|            |—— auth.proto
|            |—— auth_grpc.pb.go
|        |—— services
|            |—— auth.go
|        |—— utils
|            |—— hash.go
|            |—— jwt.go
|—— go-grpc-order-svc
|    |—— Makefile
|    |—— cmd
|        |—— main.go
|    |—— go.mod
|    |—— go.sum
|    |—— pkg
|        |—— client
|            |—— product_client.go
|        |—— config
|            |—— config.go
|            |—— envs
|                |—— dev.env
|        |—— db
|            |—— db.go
|        |—— models
|            |—— order.go
|        |—— pb
|            |—— order.pb.go
|            |—— order.proto
|            |—— order_grpc.pb.go
|            |—— product.pb.go
|            |—— product.proto
|            |—— product_grpc.pb.go
|        |—— services
|            |—— order.go
|—— go-grpc-product-svc
|    |—— .gitignore
|    |—— Makefile
|    |—— cmd
|        |—— main.go
|    |—— go.mod
|    |—— go.sum
|    |—— notes.txt
|    |—— pkg
|        |—— config
|            |—— config.go
|            |—— envs
|                |—— dev.env
|        |—— db
|            |—— db.go
|        |—— models
|            |—— product.go
|            |—— stock_decrease_log.go
|        |—— pb
|            |—— product.pb.go
|            |—— product.proto
|            |—— product_grpc.pb.go
|        |—— services
|            |—— product.go
|—— readme.MD
```
## Code Details


### Tested Platform
- software
  ```
  OS: Ubuntu 
  Go: 1.18.1 
  ```
