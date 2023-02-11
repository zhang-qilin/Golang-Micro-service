# gRPC教程



## 什么是微服务

### 单体架构

![image-20230211100317957](README.assets/image-20230211100317957.png)

一些致命缺点：

1. 一旦某个服务宕机，会引起整个应用不可用，隔离性差
2. 只能整体应用进行伸缩，浪费资源，可伸缩性差
3. 代码耦合在一起，可维护性差

### 微服务架构

为解决单体架构的问题，这就需要将服务进行拆分，单独管理和维护

![image-20230211095814091](README.assets/image-20230211095814091.png)

以上的架构，解决了单体架构的弊端

但是，同时也是引入了新的问题：

1. 代码冗余
2. 服务与服务之间存在调用关系

#### 代码冗余问题

服务未拆解之前，公告的功能有统一的实现，比如：认证、授权、限流等，但是服务拆解之后，每一个服务可能都要实现一遍

为此出现了以下的解决方案：

1. 由于为了保存对外提供的服务一致性，引入了网关的概念，由网关根据不同的请求，将其转发到不同的服务(路由功能)，由于入口的一致性，可以在网关实现公共的一些功能。    
2. 可以将公共的功能抽取出来，形成一个新的服务，比如统一认证中心。

#### 服务之间调用

服务拆解之后，服务于服务之间发生的是进程和进程之间的调用，服务器和服务器之间调用。

那么就需要发起网络调用，网络调用我们能立马的想起的就是http，但是在微服务框架中，http虽然便捷方便，但是性能较低，这时候就需要引入RPC（远程过程调用），通过自定义协议发起TCP调用，来加快传输效率。

每个服务由于可能分布在成千上百的机器上，服务和服务之间的调用，会出现一些问题，比如：如何找到应该调用哪台机器上的服务，调用方可能需要维护被调用方的地址，这个地址可能很多个，增加额外的负担，这个时候就需要引入服务治理。

服务治理中有一个重要的概念，服务发现，服务发现中有一个重要的概念，注册中心。

![image-20230211102508404](README.assets/image-20230211102508404.png)

每个服务启动的时候，会将自身的服务和IP注册到注册中心，其他服务调用的时候，只需要向注册中心申请地址即可。

当然，服务和服务之间调用会发送一些问题，为了避免产生连锁的雪崩反应，引入服务容错，为了追踪一个调用所经过的服务，引入了链路追踪，等等这些就构建了一个微服务的生态。

## gRPC

gRPC是一款语言中立，平台中立，开源的远程过程调用系统，gRPC客户端和服务端可以在多种环境中允许和交互，比如Java写的服务端，可以用Go语言语言写客户端调用。

数据在进行网络传输的时候，需要进行序列化，序列化协议有很多种，比如xml、json、protobuf等…

gRPC默认使用protocol buffers，这是Google开源的一套成熟的结构数据序列化机制。

在学习gRPC之前，需要对protocol buffers有一个大致了理解。

序列化：将数据结构或对象转化成二进制串的过程。

反序列化：将在序列化过程中产生的二进制串转换成数据结构或对象的进程。

## protobuf

protobuf是Google开源的一种数据格式，适合高性能，对影响速度有要求的数据传输场景。因为protobuf是二进制数据格式，需要编码和解码。数据本身不具有可读性。因此只能反序列化之后得到的真正的可读的数据。

优势：

1. 序列化后体积相比JSON和XML很小，适合网络传输
2. 支持跨平台多语言
3. 消息格式升级和兼容性还不错
4. 序列化反序列化很快

劣势：

这个目前俺也不太清楚



### 安装

第一步：下载通用编译器

地址：[Releases · protocolbuffers/protobuf (github.com)](https://github.com/protocolbuffers/protobuf/releases)

根据不同的操作系统，下载不同的包

第二步：配置环境变量

第三步：安装Go专用的protoc的生成器

```bash
go get github.com/golang/protobuf/protoc-gen-go
```

如何使用protobuf：

1. 定义一种源文件，扩展名为`.proto`，使用这种源文件，可以定义存储类的内容(消息类型、格式)
2. protobuf有自己的编译器`protoc`，可以将`.proto`编译成对应语言的文件，就可以进行使用了

### hello world

假设现在需要传输用户的消息，其中有username和age两个字段

```protobuf
// 指点当前proto语法版本，有2和3版本
syntax = "proto3";
// option go_package = "path;name";   // path表示生成的go文件的存放位置，会自动生成目录  name表示生成的go文件所属的包名
option go_package = ".;./service";
package service;

massage User {
	string username = 1;
	int32 age = 2;
}
```

### proto文件    

#### message介绍

`message`：`protobuf`中定义了一个消息类型是通过关键字`message`字段指定的。

消息就是需要传输的数据格式的定义。

message关键字类似于C++ 中的class，Java中的class，go中的struct

例如：

```protobuf
message User {
	string username = 1;
	int32 age = 2;
}
```

在消息中承载的数据分别对应于每一个字段。

其中每个字段都有一个名字和一种类型。

#### 字段规则

`required`：消息体中必填字段，不设置会导致编码异常。

`optional`：消息体中可选字段。

`reqeated`：消息体中可重复字段，重复的值会顺序被保留，在go中重复的会被定义为切片。

```protobuf
message User{
    string username = 1;
    int32 age = 2;
    optional string password = 3;
    repeated string addresses = 4;
}
```

#### 字段映射

| .proto   | Notes                                                        | Go      | Ruby                 | C++    | Java       | Python         | C#         |
| -------- | ------------------------------------------------------------ | ------- | -------------------- | ------ | ---------- | -------------- | ---------- |
| double   |                                                              | float64 | Float                | double | double     | float          | double     |
| float    |                                                              | float32 | Float                | float  | float      | float          | float      |
| int32    | 使用变长编码对于负值的效力很低，如果你的域有，可能会有负值，请使用sint64代替 | int32   | Fixnum or Bignum     | int32  | int        | int            | int        |
| int64    |                                                              | int64   | Bignum               | int64  | long       | ing/long[3]    | long       |
| uint32   | 使用变长编码                                                 | uint32  | Fixnum or Bignum     | uint32 | int[1]     | int/long[3]    | uint       |
| uint64   | 使用变长编码                                                 | uint64  | Bignum               | uint64 | long[1]    | int/long[3]    | ulong      |
| sint32   | 使用变长编码，这些编码在负值时比int32高效得多                | int32   | Fixnum or Bignum     | int32  | int        | intj           | int        |
| sint64   | 使用变长编码，有符号的整形值。编码时通常得比int64高          | int64   | Bignum               | int64  | long       | int/long[3]    | long       |
| fixed32  | 总是4个字节，如果数值总是比总比228大的话，这个类型会比uint32高效 | uint32  | Fixnum or Bignum     | uint32 | int[1]     | int            | uint       |
| fixed64  | 总是8个字节，如果数值总是比总比256大的话，这个类型会比uint64高效 | uint64  | Bignum               | uint64 | long[1]    | int/long[3]    | ulong      |
| sfixed32 | 总是4个字节                                                  | int32   | Fixnum or Bignum     | int32  | int        | int            | int        |
| sfixed64 | 总是8个字节                                                  | int64   | Bignum               | int64  | long       | int/long[3]    | long       |
| bool     |                                                              | bool    | TrueClass/FalseClass | bool   | boolean    | boolean        | bool       |
| string   | 一个字符串必须是UTF-8编码或者7 - bit ASCII编码的文本         | string  | String(UTF-8)        | string | String     | str/unicode[4] | string     |
| bytes    | 可能包含任何顺序的字节数据                                   | []byte  | String(ASCII-8BIT)   | string | ByteString | str            | ByteString |

#### 默认值

protobuf3删除了protobuf2中用来设置默认值的default关键字，取而代之的是protobuf3为各类型定义的默认值，也就是约定的默认值，如下所示：

| 类型         | 默认值                                                       |
| ------------ | ------------------------------------------------------------ |
| bool         | false                                                        |
| 整形         | 0                                                            |
| string       | 空字符串“”                                                   |
| 枚举类型enum | 第一个枚举元素值，因为Protobuf3强制要求第一个枚举元素的值必须是0，所以枚举的默认值就是0 |
| message      | 不是nill，而是DEFAULT_INSTANCE                               |

#### 标识号

`标识符`：在消息体的定义中，每个字段都必须要有唯一的标识号，标识号是[0,2^29-1]范围内的一个整数

```protobuf
message Person {
	string name = 1;
	int32 id = 2;
	optional string email = 3;
	repeated string phones = 4;
}
```

以Person为例，name=1，id=2，email=3，phpnes=4中1-4就是标识号。

#### 定义多个消息类型

一个proto文件可以定义多个消息类型

```protobuf
message UserRequest {
	string name = 1;	
	int32 id = 2;
	optional string email = 3;
	repeated string phones = 4;
}

message UserResponse {
	string name = 1;				// (位置1)
	int32 id = 2;
	optional string email = 3;
	repeated string phones = 4;		// (位置4)
}
```



#### 嵌套消息

#### 服务定义(Service)

## gRPC实例

### RPC和gRPC介绍



#### 