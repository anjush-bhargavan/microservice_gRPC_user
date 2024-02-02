package interfaces

import (
	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/entities"
	pb "github.com/anjush-bhargavan/microservice_gRPC_user/pkg/pb"
)

type UserServiceInter interface {
	SignupService(userpb *pb.UserCreate) (*entities.User, error)
	LoginService(userpb *pb.LoginRequest) (*entities.User, error)
	FindAllUsersService(p *pb.FetchAll) ([]*entities.User, error)
	// CreateUserService(p *pb.UserCreate) (*pb.CommonResponse, error)
	EditUserService(p *pb.UserCreate) (*pb.UserCreate,error) 
	FindUserByEmailService(p *pb.Email) (*pb.UserDetail,error)
	FindUserByIDService(p *pb.ID) (*pb.UserDetail,error)
	DeleteUserService(p *pb.ID) (*pb.CommonResponse,error)
}
