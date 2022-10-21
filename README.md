# golang-grpc-poc

## Introduction
A poc TOPK element project to test gRPC streaming, bi-direction streaming, standard

## Prerequisite
- Install protobuf compliter 3.x -> https://github.com/protocolbuffers/protobuf/
- Install golang 1.19+

## How to Run
```
// navigate to dir ./simple

// Start a grpc server
go run server/main.go

// Start client
// Open another terminal

// test standard, will auto add 5 key to the heap, where the max size of heap is 5
go run client/main.go  -race --f 1

// test standard, will increase the prioity of the item in the heap if exist else will add item to the heap and push item out if exceed max size
go run client/main.go  -race --f 1 --key key1

// test standard, will return the heap in prioity
go run client/main.go  -race --f 2

// test standard, will streamming return the item in heap (un order)
go run client/main.go  -race --f 3

// test standard, will streamming add item to heap from client, and return all item in heap (un order) from streamming
go run client/main.go  -race --f 4
```
