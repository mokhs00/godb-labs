package repository

import "godb-labs/frameworks/entity"

type UserRepository interface {
	// FindUserByID finds a user by ID
	FindUserByID(id int) (*entity.User, error)
	// FindUserByEmail finds a user by email
	FindUserByEmail(email string) (*entity.User, error)
}
