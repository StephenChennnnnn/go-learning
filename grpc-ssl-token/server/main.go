package main

import (
	"context"
	"fmt"
	"go-learning/grpc-ssl-token/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
)

func main() {
	// TLS
	creds, err := credentials.NewServerTLSFromFile(
		"./grpc-ssl-token/key/server.pem",
		"./grpc-ssl-token/key/server.key")
	if err != nil {
		grpclog.Fatal("load creds file failed")
	}

	// instance grpc server, open TLS
	server := grpc.NewServer(grpc.Creds(creds),
		grpc.UnaryInterceptor(TokenInterceptor))

	message.RegisterMathServiceServer(server, new(MathManager))

	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}

type MathManager struct {
}

func (mm *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {
	// thought meatdata
	//md, exist := metadata.FromIncomingContext(ctx)
	//if !exist {
	//	return nil, status.Errorf(codes.Unauthenticated, "no token info")
	//}
	//
	//var appKey string
	//var appSecret string
	//
	//if key, ok := md["appid"]; ok {
	//	appKey = key[0]
	//}
	//
	//if secret, ok := md["appkey"]; ok {
	//	appSecret = secret[0]
	//}
	//
	//if appKey != "hello" || appSecret != "20200623" {
	//	return nil, status.Errorf(codes.Unauthenticated, "token unlow")
	//}

	fmt.Println(" server add method")
	result := request.Args1 + request.Args2
	fmt.Println("result: ", result)
	response = new(message.Response)
	response.Code = 1
	response.Message = "succeed"
	return response, nil
}

// unaryServerInterceptor
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// thought metadata
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "no token info")
	}

	var appKey string
	var appSecret string
	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}
	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "hello" || appSecret != "20200623" {
		return nil, status.Errorf(codes.Unauthenticated, "token unlow")
	}
	// continue process request
	return handler(ctx, req)
}
