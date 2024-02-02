package service

import (
	"errors"
	"fmt"

	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/entities"
	pb "github.com/anjush-bhargavan/microservice_gRPC_user/pkg/pb"
	interfaces "github.com/anjush-bhargavan/microservice_gRPC_user/pkg/repo/interface"
	inter "github.com/anjush-bhargavan/microservice_gRPC_user/pkg/service/interface"
)

type UserService struct {
	repo interfaces.UserRepoInter
}

func NewUserService(repo interfaces.UserRepoInter) inter.UserServiceInter {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) SignupService(userpb *pb.UserCreate) (*entities.User, error) {
	user := &entities.User{
		UserName: userpb.Username,
		Email:    userpb.Email,
		Password: userpb.Password,
		Role:     "user",
	}

	err := u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *UserService) LoginService(userpb *pb.LoginRequest) (*entities.User, error) {
	user, err := u.repo.FindUserByEmail(userpb.Email)
	if err != nil {
		return nil, err
	}
	if user.Password != userpb.Password {
		return nil, errors.New("invalid password")
	}
	if user.Role != userpb.Role {
		return nil, errors.New("invalid role")
	}
	return user, nil
}

func (u *UserService) FindAllUsersService(p *pb.FetchAll) ([]*entities.User, error) {
	users, err := u.repo.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserService) EditUserService(p *pb.UserCreate) (*pb.UserCreate, error) {
	var usermodel entities.User
	usermodel.UserName = p.Username
	usermodel.Email = p.Email
	usermodel.Password = p.Password
	err := u.repo.UpdateUser(&usermodel)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (u *UserService) FindUserByEmailService(p *pb.Email) (*pb.UserDetail, error) {
	user, err := u.repo.FindUserByEmail(p.Email)
	if err != nil {
		return nil, err
	}
	return &pb.UserDetail{
		Id:       uint32(user.ID),
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

func (u *UserService) FindUserByIDService(p *pb.ID) (*pb.UserDetail, error) {
	user, err := u.repo.FindUserByID(uint(p.Id))
	if err != nil {
		fmt.Println("psoioi")
		fmt.Println(err.Error())
		return nil, err
	}
	return &pb.UserDetail{
		Id:       uint32(user.ID),
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

func (u *UserService) DeleteUserService(p *pb.ID) (*pb.CommonResponse, error) {
	err := u.repo.DeleteUser(uint(p.Id))
	if err != nil {
		return nil, err
	}
	return &pb.CommonResponse{
		Status:  "success",
		Error:   "",
		Message: "user deleted ",
	}, nil
}
