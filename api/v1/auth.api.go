package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/raminfp/backend_gin/serilizers"
	"github.com/raminfp/backend_gin/services"
	"net/http"
)

type AuthAPI interface {
	Register(context *gin.Context)
}

type authAPI struct {
	authService services.AuthService
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
