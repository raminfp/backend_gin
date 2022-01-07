package repository

import (
	"gorm.io/gorm"
)

type UserRepository interface {
}

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &userRepository{
		conn: connection,
	}
}
