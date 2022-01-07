package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/raminfp/backend_gin/entity"
	"github.com/raminfp/backend_gin/pkg/jwt"
	"github.com/raminfp/backend_gin/serilizers"
	"github.com/raminfp/backend_gin/services"
	"net/http"
)

type AuthAPI interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
}

type authAPI struct {
	authService services.AuthService
}

func (a authAPI) Login(context *gin.Context) {
	var loginRequest serilizers.LoginRequest
	var user entity.User
	err := context.ShouldBind(&loginRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	result, err := a.authService.LoginVerify(loginRequest.Email, loginRequest.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	if result == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	user.Email = loginRequest.Email
	jwt := jwt.Jwt{}
	token, err := jwt.CreateToken(user)
	context.JSON(http.StatusOK, gin.H{"data": token, "status": true})
	return
}

func NewAuthAPI(authService services.AuthService) AuthAPI {
	return &authAPI{
		authService: authService,
	}
}

func Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (a authAPI) Register(context *gin.Context) {
	var registerRequest serilizers.RegisterRequest
	err := context.ShouldBind(&registerRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	user, err := a.authService.AddUserService(registerRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": false})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user, "status": true})
	return
}
