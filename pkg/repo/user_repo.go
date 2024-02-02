package repo

import (
	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/entities"
	"gorm.io/gorm"
)

// UserRepository for connecting db to UserRepoInter methods
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepo creates an instance of user repo
func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser creates a newuser in database else returns error
func (u *UserRepository) CreateUser(user *entities.User) error {
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindUserByID finds the user in the database using ID
func (u *UserRepository) FindUserByID(id uint) (*entities.User, error) {
	var user *entities.User
	err := u.DB.First(&user, id).Error
	return user, err
}

// UpdateUser updates user details in the databse
func (u *UserRepository) UpdateUser(user *entities.User) error {
	err := u.DB.Save(&user).Error
	return err
}

// DeleteUser deletes a user using the id
func (u *UserRepository) DeleteUser(id uint) error {
	err := u.DB.Delete(&entities.User{}, id).Error
	return err
}

// FindUserByEmail finds the user by email
func (u *UserRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user *entities.User
	err := u.DB.Where("email = ?",email).First(&user,).Error
	return user, err
}

//FindAllUsers finds the all the user details in the database
func (u *UserRepository) FindAllUsers() ([]*entities.User, error) {
	var users []*entities.User
	err := u.DB.Find(&users).Error
	return users, err
}
