package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-learning/protobuf/example"
	"os"
)

func main() {
	msg_test := &example.Person{
		Name: proto.String("cxy"),
		Age:  proto.Int32(18),
		From: proto.String("china"),
	}

	// 序列化
	msgDataEncoding, err := proto.Marshal(msg_test)
	if err != nil {
		panic(err.Error())
		return
	}

	msgEntity := example.Person{}
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	fmt.Printf("name: %s\n", msgEntity.GetName())
	fmt.Printf("name: %v\n", *msgEntity.Name)
	fmt.Printf("age: %d\n", msgEntity.GetAge())
	fmt.Printf("from: %s\n", msgEntity.GetFrom())
}
