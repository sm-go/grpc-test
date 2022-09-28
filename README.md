# gRPC TESTING

`
g - Google (Golang)
R - Remote to other micro services
P - Procedure (Golang is a procedural programming language)
C - Calls 

gRPC - free, open-source (Google) - Cloud Native Computation Foundation(CNCF) - like Docker, Kubernetes

Default - HTTP V.2.0

gRPC - function - to define REQUEST and RESPONSE for RPC & handles all the rest for you
(fast and efficient, low latency, support for streaming, language independent and super easy to plug in Authentication, load balancing, logging and monitoring)

- need to define the messages and services using Protocol Buffers

one .proto file works for over 12 programming language (server/client)

- is a future of micro-services API, and mobile-server API (may be WEB API) 

The payload is binary (very efficient to send/receive on a network and serialize/de-serializer on a CPU)

http/1.1 - header is a plain text

http/2.0 - the clients & server can push message in parallel over the same TCP connection (2015)
(this greatly reduces latency)
 - REQUEST တစ်ကြိမ်လေး လုပ်ရုံလေးနဲ့ SERVER က Messages တွေအများကြီးကို Push လုပ်ပေးနိုင်တယ်။
 - HEADER ကိုလည်း Compressed လုပ်ပေးတယ်။ is Binary. 
  - is Secure. (SSL is not required but recommended by default)
 - gRPC is SSL built-in

 4 Types of gRPC
 ===============
1. Unary 			- Client <=> Server | (1 <===> 1 ) 
2. Server Streaming	- Client <=> Server | (1 <===> many)
3. Client Streaming	- Client <=> Server | (many <===> 1)
4. Bi Dir Streaming	- Client <=> Server | (many <===> many)
keyword for straming is 'stream'

gRPC is 25 times more performant than REST API

install
======

====>> GO-ZERO.DEV Website 

$ go get -u google.golang.org/grpc

$ go get -u github.com/golang/protobuf/protoc-gen-go

$ protoc

$ protoc path/to/file.proto --go_out=plugins=grpc:.
(#!/bin/bash) -> file.sh

$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative personpb/person.proto

 - service PersonService {} မပါခဲ့ဘူးဆိုလျှင် person_grpc.pb.go file ထွက်လိမ့်မည် မဟုတ်ပါ။

step by step
===========
in project folder
- server
   - main.go
- client
   - main.go
- servicepb
   - service.proto
   - service_gen.go
   - service_pb.go 
`