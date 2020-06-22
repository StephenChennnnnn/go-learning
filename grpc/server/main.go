package main

import (
	"context"
	"errors"
	"fmt"
	"go-learning/grpc/message"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {
	server := grpc.NewServer()

	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}

type OrderServiceImpl struct {
}

func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {
	orderMap := map[string]message.OrderInfo{
		"202006200001": {OrderId: "202006200001", OrderName: "clothes", OrderStatus: "yes"},
		"202006210001": {OrderId: "202006210001", OrderName: "tips", OrderStatus: "yes"},
		"202006210002": {OrderId: "202006210002", OrderName: "food", OrderStatus: "no"},
	}

	var response *message.OrderInfo
	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "null"}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			fmt.Println(result)
			return &result, nil
		} else {
			return nil, errors.New("server error")
		}
	}
	return response, nil
}
