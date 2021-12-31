package services

import (
	"github.com/raminfp/backend_gin/entity"
	"github.com/raminfp/backend_gin/repository"
	"github.com/raminfp/backend_gin/serilizers"
)

type AuthService interface {
	AddUserService(registerRequest serilizers.RegisterRequest) (string, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

func (a authService) AddUserService(registerRequest serilizers.RegisterRequest) (string, error) {
	user := entity.User{}
	user.Email = registerRequest.Email
	user.Firstname = registerRequest.Firstname
	user.Lastname = registerRequest.Lastname
	user.Password = registerRequest.Password

	addUser, err := a.authRepository.AddUser(user)
	if err != nil {
		return "", err
	}
	return addUser, nil
}
