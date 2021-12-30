package middleware

import "github.com/gin-gonic/gin"

func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Method not allowed"})
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Route not defined"})
	}
}
