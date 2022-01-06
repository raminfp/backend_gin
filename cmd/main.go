package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raminfp/backend_gin/routes"
)

func main() {

	r := routes.Urls()
	gin.SetMode(gin.ReleaseMode)
	err := r.Run()
	if err != nil {
		return
	}
}
