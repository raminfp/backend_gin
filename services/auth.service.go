package services

import (
	"errors"
	"github.com/mashingan/smapping"
	"github.com/raminfp/backend_gin/entity"
	"github.com/raminfp/backend_gin/repository"
	"github.com/raminfp/backend_gin/serilizers"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	AddUserService(registerRequest serilizers.RegisterRequest) (string, error)
	LoginVerify(email string, password string) (string, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func (a authService) LoginVerify(email string, password string) (string, error) {
	user, err := a.authRepository.FindByEmail(email)
	if err != nil {
		return "error", err
	}
	isValidPassword := comparePasswords(user, []byte(password))
	if !isValidPassword {
		errors.New("failed to login, because password is not matched")
	}
	return "Ok", nil

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

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
