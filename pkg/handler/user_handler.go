package handler

import (
	"context"
	"fmt"

	pb "github.com/anjush-bhargavan/microservice_gRPC_user/pkg/pb"
	inter "github.com/anjush-bhargavan/microservice_gRPC_user/pkg/service/interface"
)

type UserHandler struct {
	pb.UserServiceServer
	svc inter.UserServiceInter
}

func NewUserHandler(svc inter.UserServiceInter) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) UserSignup(ctx context.Context,p *pb.UserCreate) (*pb.Response, error) {
	result, err := u.svc.SignupService(p)
	if err != nil {
		return &pb.Response{
			Status:  "failed",
			Error:   "user signup error",
			Message: "",
		}, err
	}
	msg := fmt.Sprintf("User created %v", result)
	return &pb.Response{
		Status:  "Success",
		Error:   "",
		Message: msg,
	}, nil
}

func (u *UserHandler) UserLogin(ctx context.Context,p *pb.LoginRequest) (*pb.Response, error) {
	result, err := u.svc.LoginService(p)
	if err != nil {
		msg := fmt.Sprintf("error: %v", err)
		return &pb.Response{
			Status:  "Failed",
			Error:   "Login error",
			Message: msg,
		}, err
	}

	msg := fmt.Sprintf("Logged in : %v", result)
	return &pb.Response{
		Status:  "success",
		Error:   "",
		Message: msg,
	}, nil
}
