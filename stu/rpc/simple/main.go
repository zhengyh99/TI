package main

import (
	"TI/stu/rpc/simple/example"

	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	msg := &example.Person{
		Name: proto.String("Vick"),
		Age:  proto.Int32(20),
		From: proto.String("KOR"),
	}
	//序列化
	enMsg, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("proto marshal error:", err)
	}
	deMsg := &example.Person{}
	//反序列化
	err = proto.Unmarshal(enMsg, deMsg)
	if err != nil {
		fmt.Println("proto unmarshal error:", err)
	}
	fmt.Printf("Name : %s \n\n", deMsg.GetName())
	fmt.Printf("age : %d \n\n", deMsg.GetAge())
	fmt.Printf("from : %s \n\n", deMsg.GetFrom())

}
