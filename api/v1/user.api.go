package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/raminfp/backend_gin/services"
	"net/http"
)

type UserAPI interface {
	Me(context *gin.Context)
}

type userAPI struct {
	userService services.UserService
}

func (u userAPI) Me(context *gin.Context) {

	userEmail, err := context.Get("userEmail")
	if !err {
		context.JSON(http.StatusBadRequest, gin.H{"error": "email is invalid", "status": false})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": userEmail, "status": true})
	return
}

func NewUserAPI(userService services.UserService) UserAPI {
	return &userAPI{
		userService: userService,
	}
}

func (u userAPI) Register(context *gin.Context) {
	panic("implement me")
}
