package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/raminfp/backend_gin/pkg/jwt"
	"net/http"
)

func AthorizationJWT(jwtService jwt.JwtService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authToken := context.GetHeader("Athorization")
		if authToken == "" {
			//context.JSON(http.StatusBadRequest, gin.H{"error": "Missing token", "status": false})
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing token", "status": false})
			return
		}
		user, err := jwtService.ValidateToken(authToken)
		if err != nil {
			//context.JSON(http.StatusBadRequest, gin.H{"error": "Missing token", "status": false})
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token", "status": false})
			return
		}
		context.Set("userEmail", user.Email)
		return
	}
}
