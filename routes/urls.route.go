package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/raminfp/backend_gin/api/v1"
	"github.com/raminfp/backend_gin/middleware"
)

func Urls() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.NoRoute(middleware.NoRouteHandler())
	r.HandleMethodNotAllowed = true
	r.NoMethod(middleware.NoMethodHandler())

	apiV1 := r.Group("api/v1")
	{
		auth := apiV1.Group("auth")
		{
			auth.GET("/:name", v1.Index)
		}
		user := apiV1.Group("user")
		{
			user.POST("/home", v1.Home)
			user.POST("/login", v1.Login)
			//user.POST("/home", v1.Me)

		}
	}
	return r
}
