# GRPC

## GRPC介绍

### 微服务

#### 单体架构

1. 一旦某个服务宕机，会引起整个应用不可用，隔离性差
2. 只能整体应用进行伸缩，浪费资源，可伸缩性差
3. 代码辗合在一起，可维护性差

微服务架构：解决了单体架构的弊端

但同时引入了新的问题

1. 代码冗余
2. 服务和服务之间存在调用关系

服务拆分后，服务和服务之间发生的是进程和进程之间的调用，服务器和服务器之间的调用。

那么就需要发起网络调用，网络调用我们能立马想起的就是http，但是在微服务架构中，http虽然便捷方便，但性能较低，这时候就需要引入
RPC(远程过程调用)，通过自定义协议发起TCP调用，来加快传输效率。

[gRPC官方文档](https://grpc.io/)

[gRPC 官方文档中文版_V1.0 (oschina.net)](http://doc.oschina.net/grpc)

RPC的全称是Remote Procedure Call ，远程过程调用，这是一种协议，是用来屏蔽分布式计算中的各种调用细节，使得你可以像是本地调用一样直接调用一个远程的函数。

客户端与服务端沟通的过程

1. 客户端 发送 数据 (以字节流的方式)
2. 服务端 接收并且解析，根据约定 知道要执行什么。然后把结果返回给客户

RPC：

1. RPC就是将上述过程封装下，使得操作更加优化
2. 使用一些大家都认可得协议，使其规范化
3. 做成一些框架，直接或间接产生利益

而gRPC又是什么呢？用官方得话来解释：

A high performance, open source universal RPC framework.

一个高性能、开源的通用RPC框架

在gRPC中，我们称调用方为client,被调用方为server。跟其他的RPC框架一样，gRPC也是基于"服务定义"的思想。简单的来讲，就是我们通过
某种方式来描述一个服务，这种描述方式是语言无关的。在这个“服务定义“的过程中，我们描述了我们提供的服务服务名是什么，有哪些方法可以
被调用，这些方法有什么样的入参，有什么样的回参。

也就是说，在定义好了这些服务、这些方法之后，gRPC会屏蔽底层的细节，client只需要直接调用定义好的方法，就能拿到预期的返回结果。对于server端来说，还需要实现我们定义的方法。同样的，gRPC也会帮我们屏蔽底层的细节，我们只需要实现所定义的方法的具体逻辑即可。

你可以发现，在上面的描述过程中，所谓的"服务定义“，就跟定义接口的语义是很接近的。我更愿意理解为这是一种"约定“，双方约定好接口，然后server实现这个接口，client调用这个接口的代理对象。至于其他的细节，交给gRPC。

此外，gRPC还是语言无关的。你可以用C++作为服务端，使用Golang、Java等作为客户端。为了实现这一点，我们在"定义服务“和在编码和解码的过程中，应该是做到语言无关的。

![Concept Diagram](README.assets/landing-2.svg)

因此gRPC使用了Protocol Buffss，这是谷歌开源得一套成熟得数据结构序列化机制。

你可以把他当成一个代码生成工具以及序列化工具。这个工具可以把我们定义的方法，转换成特定语言的代码。比如你定义了一种类型的参数，他会帮你转换成Golang中的struct结构体，你定义的方法，他会帮你转换成func函数。此外，在发送请求和接受响应的时候，这个工具还会完成对应的编码和解码工作，将你即将发送的数据编码成gRPC能够传输的形式，又或者将即将接收到的数据解码为编程语言能够理解的数据格式。

序列化：将数据结构或对象转换成二进制串的过程

反序列化：将在序列化过程中所产生的二进制串转换成数据结构或者对象的过程

protobuf是谷歌开源的一种数据格式，适合高性能，对响应速度有要求的数据传输场景。因为protobuf是二进制数据格式，需要编码和解码。数据本身不具有可读性。因此只能反序列化之后得到真正可读的数据。

优势

1. 序列化后体积相比JSON和XML很小，适合网络传输
2. 支持跨平台多语言
3. 消息格式升级和兼容性还不错
4. 序列化反序列化速度很快

## 安装Protobuf

### 下载[protocol buffers](https://github.com/protocolbuffers/protobuf/releases)

Protocol buffers，通常称为Protobuf，是Google开发得一种协议，用于允许对结构化数据进行序列化和反序列化。它在开发程序以通过网络相互通信或存储数据时很有用。谷歌开发它得目的是提高了一种比XML和JSON更加方便得方式来通信。

我们将找到所有系统得所有zip文件，基于我们的操作系统位版本( 64位或 32 位)，来下载指定的版本，解压压缩包后将bin目录添加到环境变量中即可

### 安装gRPC的核心库

```bash
go get google.golang.org/grpc
```

上面安装的是protocol编译器。它可以生成各种不同的代码。因此，除了这个编译器，我们还需要配合各个语言代码的生成工具，对于Golang来说，称为protoc-gen-go。不过这里有个小小的坑，github.com/golang/protobuf/protoc-gen-go和google.golang.org/protobuf/cmd/protoc是不同的，区别在于前者是旧版本，后者是google接管后的新版本，他们之间的API是不同的，也就是说用来生成的命令，以及生成的文件都是不一样的。因为目前的gRPC-go源码中的example用的是后者的生成方式，为了与时俱进，我们也采用最新的方式，你需要安装两个库：

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

执行以上两句命令后会在GoPATH下的bin目录下生成对应的可执行文件

## Proto文件编写

```protobuf
// 这是在说明我们使用的是proto3语法
syntax = "proto3";

// 这部分的内容是关于最后生成的go文件是处于哪个包中， .代表当前目录生成， service代表了生成的go文件的包名是service。
option go_package = ".;service";

// 然后需要定义一个服务，在这个服务中需要有一个方法，这个方法可以接受客户端的参数，再放回服务端的响应。
// 其实很容易可以看出，我们定义了一个server，称为SayHello，这个服务中有一个rpc方法，名为SayHello。
service SayHello {
  rpc SayHello(HelloRequest) returns (HelloResponse){}
}

// message 关键字，其实可以理解为Golang中的结构体。
// 这里比较特别的是变量后面的“赋值”，注意，这里并不是赋值，而是再定义这个变量在这个message中的位置。
message HelloRequest {
  string requestName = 1;
  int64 age = 2;
}

message HelloResponse {
  string responseMsg = 1;
}
```

在编写完上面的内容后，在helloword/proto目录下执行以下的命令：

```bash
protoc --go_out=. hello.proto
protoc --go-grpc_out=. hello.proto
```

## Proto文件介绍

### message

message：protobuf中定义一个消息类型式是通过关键字message字段指定的，消息就是需要传输的数据格式的定义。	

message关键字类似于C++ 中的class，Java中的class，go中的struct。

在消息中承载的数据分别对应于每一个字段，其中每个字段都有一个名字和类型。

一个proto文件中可以定义多个消息类型。

### 字段规则

required：消息体中必填的字段，不设置会导致编码异常。在protobuf2中使用，在protobuf3中被删去。

optionl：消息体中可选字段。protobuf3没了required，optional等说明关键字，都默认为optional。

repeated：消息体中重复字段，重复的值的顺序会被保留在go中重复的会被定义为切片。

### 消息号

在消息的定义中，每个字段都必须有一个唯一的标识号，标识号是[1, 2^29-1]范围内的一个整数。

### 嵌套消息

可以在其它消息类型中定义、使用消息类型，在下面的例子中，person消息会定义在PersonInfo消息内

```protobuf
message PersionInfo {
	message Person {
		string name = 1;
		int32 heigth = 3;
		repeated int32 weight = 3;
	}
	repeated Person info = 1;
}
```

在父消息类型的外部调用这个消息类型，需要PersonInfo.Person的形式使用它，如下所示：

```protobuf
message PersonMessage {
	PersonIon.Person info = 1;
}
```

### 服务定义

如果想将消息类型作用在RPC系统中，可以在. proto文件中定义一个RPC服务接口，protocol buffer编译器将会根据所选择的不同语言生成服务接口代码以及存根。

```protobuf
service SearchService {
	// rpc 服务函数名(参数) 返回 (返回参数)
	rpc Search(SearchRequest) returns (SearchResponse)
}
```

上诉代表表示，定义了一个RPC服务，该方法接受SearchRequest返回SearchResponse

## 服务端编写

- 创建gRPC Server对象，可以理解为它是Server端的抽象对象

- 将Server（其包含需要被调用的服务端端口）注册到gRPC Server的内部注册中心。

  这样在接受到请求时，通过内部的服务发现，发现该服务端口并转接进行逻辑处理

- 创建Listen，监听TCP端口

- gRPC Server开始lit.Accept，直到Stop



## 客户端编写

- 创建与给定目标（服务端）的了解交互

- 传教Server的客户端对象

- 发送RPC请求，等待同步响应，得到回调后返回响应结果

- 输出响应结果
