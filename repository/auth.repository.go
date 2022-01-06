package repository

import (
	"github.com/raminfp/backend_gin/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type AuthRepository interface {
	AddUser(user entity.User) (string, error)
	FindByEmail(email string) (string, error)
	FindByID(userId int64) (string, error)
}

type authRepository struct {
	conn *gorm.DB
}

func (a authRepository) FindByID(userId int64) (string, error) {
	var user entity.User
	res := a.conn.Where("id = ?", userId).Take(&user)
	if res.Error != nil {
		return "", res.Error
	}
	return "founded", nil
}

func (a authRepository) FindByEmail(email string) (string, error) {
	var user entity.User
	res := a.conn.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return "", res.Error
	}
	return "founded", nil
}

func NewAuthRepository(connection *gorm.DB) AuthRepository {
	return &authRepository{
		conn: connection,
	}
}

func (a authRepository) AddUser(user entity.User) (string, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	}
	a.conn.Save(&user)
	return "insert", nil
}

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
