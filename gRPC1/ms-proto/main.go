/*
* @Time ： 2023-02-11 16:43
* @Auth ： 张齐林
* @File ：main.go
* @IDE ：GoLand
 */
package main

import (
	"fmt"
	"ms-proto/service"
	
	"google.golang.org/protobuf/proto"
)

func main() {
	user :=&service.User{
		Username: "张齐林",
		Age: 18,
	}
	// 序列化的过程
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	
	// 反序列化的过程
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(newUser.String())
}
