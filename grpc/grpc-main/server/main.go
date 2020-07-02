package main

import (
	"context"
	"errors"
	"fmt"
	"go-learning/grpc/message"
	"google.golang.org/grpc"
	"io"
	"net"
	"strconv"
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

func (os *OrderServiceImpl) GetOrderInfosCS(stream message.OrderService_GetOrderInfosCSServer) error {
	orderMap := map[string]message.OrderInfo{
		"202006200001": {OrderId: "202006200001", OrderName: "clothes", OrderStatus: "yes"},
		"202006210001": {OrderId: "202006210001", OrderName: "tips", OrderStatus: "yes"},
		"202006210002": {OrderId: "202006210002", OrderName: "food", OrderStatus: "no"},
	}
	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("server: data read end")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}

		fmt.Println(orderRequest.GetOrderId())
		result := orderMap[orderRequest.GetOrderId()]
		// send data
		err = stream.Send(&result)
		if err == io.EOF {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func (os *OrderServiceImpl) AddOrderList(stream message.OrderService_AddOrderListServer) error {
	fmt.Println("client stream RPC mode")

	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read data end")
			result := message.OrderInfo{OrderStatus: "read data end"}
			return stream.SendAndClose(&result)
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		// print receive data
		fmt.Println(orderRequest)
	}
}

func (os *OrderServiceImpl) GetOrderInfos(request *message.OrderRequest, stream message.OrderService_GetOrderInfosServer) error {
	fmt.Println("server stream RPC mode.")

	orderMap := map[string]message.OrderInfo{
		"202006200001": {OrderId: "202006200001", OrderName: "clothes", OrderStatus: "yes"},
		"202006210001": {OrderId: "202006210001", OrderName: "tips", OrderStatus: "yes"},
		"202006210002": {OrderId: "202006210002", OrderName: "food", OrderStatus: "no"},
	}
	for id, info := range orderMap {
		o, _ := strconv.Atoi(request.OrderId)
		i, _ := strconv.Atoi(id)
		if o >= i {
			fmt.Println("order id: ", id)
			fmt.Println("order info: ", info)
			// stream mode
			stream.Send(&info)
		}
	}
	return nil
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
