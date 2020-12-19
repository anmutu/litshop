package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"litshop/pb/protobuf"
	"litshop/pb/protobuf/common"
	_ "litshop/src/bootstrap"
	"litshop/src/config"
	"litshop/src/controller/authentcation"
	context2 "litshop/src/lvm/context"
	"litshop/src/lvm/literr"
	"litshop/src/pkg/logger"
	"litshop/src/request"
	"net"
)

func main() {
	var srv *grpc.Server
	srv = grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.GetString("microservice.authservice.host"), config.GetString("microservice.authservice.port")))
	if err != nil {
		logger.Fatal(err)
	}

	protobuf.RegisterAuthServiceServer(srv, &AuthService{})
	_ = srv.Serve(lis)
}

type AuthService struct {
	protobuf.UnimplementedAuthServiceServer
}

func (*AuthService) SignIn(c context.Context, in *common.Request) (*common.Response, error) {
	ctx, _ := context2.FromGrpcServer()

	r := &request.SignInRequest{}

	if in.GetSignInRequest() == nil {
		return nil, literr.NewWithCode(literr.ErrCodeInvalidRequestParams)
	}

	b, err := json.Marshal(in.GetSignInRequest())
	if err != nil {
		return nil, nil
	}
	err = json.Unmarshal(b, r)
	if err != nil {
		return nil, nil
	}

	_, err = authentcation.SignIn(ctx, r)
	return nil, err
}

func (*AuthService) SignUp(c context.Context, in *common.Request) (*common.Response, error) {
	return nil, nil
}

func (*AuthService) Logout(c context.Context, in *common.Request) (*common.Response, error) {
	return nil, nil
}
