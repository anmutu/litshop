package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "litshop/pb/protobuf/authservice"
	_ "litshop/src/bootstrap"
	"litshop/src/config"
	"litshop/src/pkg/logger"
	"net"
)

func main() {
	var srv *grpc.Server
	srv = grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.GetString("microservice.authservice.host"), config.GetString("microservice.authservice.port")))
	if err != nil {
		logger.Fatal(err)
	}

	pb.RegisterAuthServiceServer(srv, &AuthService{})
	_ = srv.Serve(lis)
}

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func (*AuthService) SignIn(context.Context, *pb.AddItemRequest) (*pb.Empty, error) {
	return nil, nil
}

func (*AuthService) SignUp(context.Context, *pb.AddItemRequest) (*pb.Empty, error) {
	return nil, nil
}

func (*AuthService) Logout(context.Context, *pb.EmptyCustomerRequest) (*pb.Empty, error) {
	return nil, nil
}
