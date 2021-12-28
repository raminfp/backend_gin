package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raminfp/backend_gin/api/v1"
	"github.com/raminfp/backend_gin/middleware"
)


func main()  {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/:name", v1.Index)
	r.POST("/home", v1.Home)
	r.POST("/login", v1.Login)


	err := r.Run()
	if err != nil {
		return
	}
}
