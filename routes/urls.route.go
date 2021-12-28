package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/raminfp/backend_gin/api/v1"
	"github.com/raminfp/backend_gin/middleware"
)

func Urls() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("/:name", v1.Index)
	}

	r.POST("/home", v1.Home)
	r.POST("/login", v1.Login)
	return r
}
