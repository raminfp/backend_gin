package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/raminfp/backend_gin/entity"
	"net/http"
)

func Login(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{"message": "OK" })
}

func Index(context *gin.Context)  {
	name := context.Param("name")
	context.JSON(http.StatusOK, gin.H{"message": name })
}

func Home(context *gin.Context)  {
	var myuser 	entity.User
	validate := validator.New()
	err := context.BindJSON(&myuser)
	err = validate.Struct(myuser)
	if err != nil {
		for _, err_user := range err.(validator.ValidationErrors) {
			context.JSON(http.StatusBadRequest, gin.H{"error": err_user.Error(), "status": false} )
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"firstname": myuser.Firstname,
		"lastname": myuser.Lastname,
		"email": myuser.Email,
	})
}
