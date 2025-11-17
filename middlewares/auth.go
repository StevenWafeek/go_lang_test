package middlewares

import (
	"OnlineServer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authraization(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
		return
	}
	userID, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
