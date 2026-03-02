package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/booking-project/utils"
)

func Authenticate(context *gin.Context) {
	// Get the token from the Authorization header
	// abortwithstatusjson is used to stop the request from being processed further and return a response immediately

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	userId, err := utils.VerifyJWTToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
