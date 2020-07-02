package main

import (
	"context"
	"fmt"
	"go-learning/grpc/message"
	"google.golang.org/grpc"
	"io"
)

func main() {
	// 1. Dial conn
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)
	//orderRequest := &message.OrderRequest{OrderId: "202006210001", TimeStamp: time.Now().Unix()}
	//orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
	//if orderInfo != nil {
	//	fmt.Println(orderInfo.GetOrderId())
	//	fmt.Println(orderInfo.GetOrderName())
	//	fmt.Println(orderInfo.GetOrderStatus())
	//}

	// server mode
	//orderInfoClient, err := orderServiceClient.GetOrderInfos(context.Background(), orderRequest)
	//if err != nil {
	//	panic(err.Error())
	//}
	//for {
	//	orderInfo, err := orderInfoClient.Recv()
	//	if err == io.EOF {
	//		fmt.Println("read end.")
	//		return
	//	}
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	fmt.Println("read info: ", orderInfo)
	//}

	// client mode
	//addOrderListClient, err := orderServiceClient.AddOrderList(context.Background())
	//if err != nil {
	//	panic(err.Error())
	//}
	//orderMap := map[string]message.OrderRequest{
	//	"202006200001": {OrderId: "202006200001", TimeStamp: time.Now().Unix()},
	//	"202006210001": {OrderId: "202006210001", TimeStamp: time.Now().Unix()},
	//	"202006210002": {OrderId: "202006210002", TimeStamp: time.Now().Unix()},
	//}
	//for _, info := range orderMap {
	//	err = addOrderListClient.Send(&info)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//}
	//for {
	//	orderInfo, err := addOrderListClient.CloseAndRecv()
	//	if err == io.EOF {
	//		fmt.Println("read data end client")
	//		return
	//	}
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(orderInfo.GetOrderStatus())
	//}

	// cs mode
	fmt.Println("cs mode")
	orderIDs := []string{"202006200001", "202006210001", "202006210002"}

	orderInfoCSClient, err := orderServiceClient.GetOrderInfosCS(context.Background())
	for _, orderID := range orderIDs {
		orderRequest := message.OrderRequest{OrderId: orderID}
		err := orderInfoCSClient.Send(&orderRequest)
		if err != nil {
			panic(err.Error())
		}
	}
	// close
	orderInfoCSClient.CloseSend()

	for {
		orderInfo, err := orderInfoCSClient.Recv()
		if err == io.EOF {
			fmt.Println("read end")
			return
		}
		if err != nil {
			return
		}
		fmt.Println("read info: ", orderInfo)
	}
}
