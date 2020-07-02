package main

import (
	"context"
	"fmt"
	"go-learning/grpc/grpc-ssl-token/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// TLS
	creds, err := credentials.NewClientTLSFromFile("./grpc-ssl-token/key/server.pem", "go-grpc-example")
	if err != nil {
		panic(err.Error())
	}
	// Token
	auth := TokenAuthentication{
		AppKey:    "hello",
		AppSecret: "20200623",
	}
	// Dail conn
	conn, err := grpc.Dial("localhost:8092",
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	serviceClient := message.NewMathServiceClient(conn)
	addArgs := message.RequestArgs{Args1: 3, Args2: 5}

	response, err := serviceClient.AddMethod(context.Background(), &addArgs)
	if err != nil {
		grpclog.Fatal(err.Error())
	}

	fmt.Println(response.GetCode(), response.GetMessage())
}

// token
type TokenAuthentication struct {
	AppKey    string
	AppSecret string
}

// organize token
func (ta *TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  ta.AppKey,
		"appkey": ta.AppSecret,
	}, nil
}

// use TLS
func (a *TokenAuthentication) RequireTransportSecurity() bool {
	return true
}
