package repository

import (
	"github.com/raminfp/backend_gin/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	AddUser(user entity.User) (string, error)
}

type authRepository struct {
	conn *gorm.DB
}

func NewAuthRepository(connection *gorm.DB) AuthRepository {
	return &authRepository{
		conn: connection,
	}
}

func (a authRepository) AddUser(user entity.User) (string, error) {
	err := a.conn.Save(&user)
	if err.Error != nil {
		return "", err.Error
	}
	return "insert", nil
}
