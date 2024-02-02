package handler

import (
	"context"
	"fmt"

	pb "github.com/anjush-bhargavan/microservice_gRPC_user/pkg/pb"
)

func (u *UserHandler) FindAllUsers(ctx context.Context, p *pb.FetchAll) (*pb.Users, error) {
	result, err := u.svc.FindAllUsersService(p)
	if err != nil {
		return nil, err
	}
	var users []*pb.UserDetail
	for _, i := range result {
		users = append(users, &pb.UserDetail{
			Username: i.UserName,
			Email:    i.Email,
			Password: i.Password,
			Role:     i.Role,
		})
	}
	response := &pb.Users{
		Userlist: users,
	}
	return response, nil
}

func (u *UserHandler) CreateUser(ctx context.Context, p *pb.UserCreate) (*pb.CommonResponse, error) {
	user := &pb.UserCreate{
		Username: p.Username,
		Email:    p.Email,
		Password: p.Password,
	}

	_, err := u.svc.SignupService(user)
	if err != nil {
		return nil, err
	}

	result := &pb.CommonResponse{
		Status:  "Success",
		Error:   "",
		Message: "user created successfully",
	}
	return result, nil
}

func (u *UserHandler) EditUser(ctx context.Context, p *pb.UserCreate) (*pb.UserCreate, error) {
	result, err := u.svc.EditUserService(p)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (u *UserHandler) DeleteUser(ctx context.Context, p *pb.ID) (*pb.CommonResponse, error) {
	res, err := u.svc.DeleteUserService(p)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserHandler) FindUserByID(ctx context.Context, p *pb.ID) (*pb.UserDetail, error) {
	user, err := u.svc.FindUserByIDService(p)
	if err != nil {
		fmt.Println("hiiiiiii")
		return nil, err
	}
	return user, nil
}

func (u *UserHandler) FindUserByEmail(ctx context.Context, p *pb.Email) (*pb.UserDetail, error) {
	user, err := u.svc.FindUserByEmailService(p)
	if err != nil {
		return nil, err
	}
	return user, nil
}
