package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"rpc/message"
	"time"
)

func main() {
	// server
	fmt.Println("start server")
	go func() {
		orderService := new(OrderService)

		rpc.Register(orderService)
		rpc.HandleHTTP()

		listen, err := net.Listen("tcp", ":8081")
		if err != nil {
			panic(err.Error())
		}
		http.Serve(listen, nil)
		select {}
	}()

	// client
	fmt.Println("start client")
	client, err := rpc.DialHTTP("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}

	timeStamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "2020061710001", TimeStamp: timeStamp}

	var response message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo", request, &response)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response)
}

type OrderService struct {
}

func (os *OrderService) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {
	orderMap := map[string]message.OrderInfo{
		"2020061700001": {OrderId: "2020061700001", OrderName: "clothes", OrderStatus: "yes"},
		"2020061710001": {OrderId: "2020061710001", OrderName: "shose", OrderStatus: "yes"},
		"2020061710002": {OrderId: "2020061710002", OrderName: "food", OrderStatus: "no"},
	}

	current := time.Now().Unix()
	fmt.Println(current)
	fmt.Println(request.TimeStamp)
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "error"}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			*response = orderMap[request.OrderId]
		} else {
			return errors.New("server error")
		}
	}
	return nil
}
