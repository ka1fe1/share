GRPC 用法简介

[TOC]

# GRPC 简介

GRPC 是 Google 开源的一种 RPC 框架

> As in many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types. On the server side, the server implements this interface and runs a gRPC server to handle client calls. On the client side, the client has a stub (referred to as just a client in some languages) that provides the same methods as the server.

GRPC 特点：

- 基于 [HTTP 2  协议](https://blog.csdn.net/hukfei/article/details/80322907)

- C/S 端都支持多语言

- 使用 Protocol Buffer 定义和序列化数据

  > 官方建议使用 proto3

GRPC 文档：

[官方文档](https://grpc.io/docs/guides/)

[中文文档](http://doc.oschina.net/grpc?t=58008)

[Protocol Buffers 3 基础教程](https://blog.csdn.net/shensky711/article/details/69696392)

---



# 基于同一个 pb 文件生成不同语言的 C、S 端，并两端实现通信

**示例说明：**

1. client 端采用 python 编写，server 端采用 go 编写

2. protocol buffer  定义文件

   > [msg_user_list.protoc](https://github.com/kaifei-bianjie/images/blob/master/msg_user_list.proto)
   >
   > [services.protoc](https://github.com/kaifei-bianjie/images/blob/master/services.proto)



**代码示例：Client 端**

生成 client 端代码：

- `python3 -m grpc_tools.protoc -I ./protocol --python_out=./generated --grpc_python_out=./generated ./protocol/msg_user_list.proto`
- `python3 -m grpc_tools.protoc -I ./protocol --python_out=./generated --grpc_python_out=./generated ./protocol/services.proto `

client 端调用代码：

```python
# -*- coding: utf-8 -*-

import grpc

from generated import services_pb2
from generated import services_pb2_grpc


def run():
    channel = grpc.insecure_channel('localhost:60051')
    stub = services_pb2_grpc.UserStub(channel)
    response = stub.GetList(services_pb2.UserListRequest(page=1, perPage=10))
    print("Get server response: \n")
    print(response)


if __name__ == '__main__':
    run()
```

**代码示例：Server 端**

生成 Server 端代码：

- `protoc -I ./protocol --go_out=plugins=grpc:./generated ./protocol/msg_user_list.proto`
- `protoc -I ./protocol --go_out=plugins=grpc:./generated ./protocol/services.proto`

Server 端实现代码：

```go
package main

import (
	"log"
	"net"
	
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "local_test/grpc/generated"
)

const (
	port = ":60051"
)

// server is used to implement UserServer.GetList.
type server struct{}

func (s server) GetList(ctx context.Context, in *pb.UserListRequest) (
	*pb.UserListResponse, error) {
	
	log.Println(ctx)
	log.Println(in)
	user1 := pb.User{
		Name: "LiLei",
		Age:  23,
	}
	user2 := pb.User{
		Name: "Zhang",
		Age:  24,
	}
	users := []*pb.User{&user1, &user2}
	
	response := pb.UserListResponse{
		User: users,
	}
	return &response, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listen on %v\n", port)
	s := grpc.NewServer()
	
	pb.RegisterUserServer(s, &server{})
	
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

```

---



# Client 端 Context 及 timeout 的用法

Client 端代码：

```python
# -*- coding: utf-8 -*-

import grpc
import sys

from generated import services_pb2
from generated import services_pb2_grpc


def run():
    channel = grpc.insecure_channel('localhost:60051')
    try:
        # wait connection
        grpc.channel_ready_future(channel).result(timeout=10)
    except grpc.FutureTimeoutError:
        sys.exit('Error connecting to server')
    else:
        stub = services_pb2_grpc.UserStub(channel)
        # set client-side timeout
        response = stub.GetList(services_pb2.UserListRequest(page=1, perPage=10), timeout=10)
        print("Get server response: \n")
        print(response)


if __name__ == '__main__':
    run()
```

[参考链接](https://blog.codeship.com/using-grpc-in-python/)