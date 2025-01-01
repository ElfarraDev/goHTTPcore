# GoHTTPCore

A lightweight, performant HTTP server implementation in Go built from scratch. This project implements core HTTP server functionality without relying on Go's net/http package, providing deep insights into server operations and HTTP protocol handling.

## Description

GoHTTPCore is an educational project that implements a HTTP server from the ground up. By managing everything from TCP sockets to HTTP request parsing, it provides hands-on experience with low-level networking concepts and the HTTP protocol.

## Features

* Custom TCP socket management for network connections
* HTTP/1.1 protocol support with request parsing
* Flexible routing system for handling different endpoints
* Response builder with status codes and header management
* Middleware chain support for request/response processing
* Static file serving capabilities
* Query parameter and form data parsing
* Basic security features and request validation

## Installation

```bash
git clone https://github.com/yourusername/gohttpcore
cd gohttpcore
go mod init gohttpcore
