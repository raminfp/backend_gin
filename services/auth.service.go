package services

import (
	"errors"
	"github.com/mashingan/smapping"
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
	userExists, _ := a.authRepository.FindByEmail(registerRequest.Email)
	if userExists != "" {
		return "", errors.New("user already exists")
	}
	//user.Email = registerRequest.Email
	//user.Firstname = registerRequest.Firstname
	//user.Lastname = registerRequest.Lastname
	//user.Password = registerRequest.Password
	mapped := smapping.MapFields(&registerRequest)
	err := smapping.FillStruct(&user, mapped)
	if err != nil {
		return "Error in mapping", nil
	}
	addUser, err := a.authRepository.AddUser(user)
	if err != nil {
		return "", err
	}
	return addUser, nil
}
