package main

import (
	"context"
	"fmt"
	"go-learning/grpc/message"
	"google.golang.org/grpc"
	"time"
)

func main() {
	// 1. Dial conn
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	orderRequest := &message.OrderRequest{OrderId: "202006210001", TimeStamp: time.Now().Unix()}
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
	if orderInfo != nil {
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	}
}
