package interfaces

import "github.com/anjush-bhargavan/microservice_gRPC_user/pkg/entities"

// UserRepo defines the methods in UserRepo
type UserRepoInter interface {
	CreateUser(user *entities.User) error
	FindUserByID(id uint) (*entities.User, error)
	UpdateUser(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
	DeleteUser(id uint) error
	FindAllUsers() ([]*entities.User, error)
}
