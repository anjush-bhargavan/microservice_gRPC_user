package entities

import "gorm.io/gorm"

// User model to stores user details
type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
	Role     string
}

// Login model for storing logging in credentials
type Login struct {
	Email    string
	Password string
}
