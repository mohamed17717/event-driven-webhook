package middlewares

import (
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		token, err := utils.JwtToken.Parse(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		var userID *uint
		userID, err = utils.JwtToken.GetUser(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		var user models.User
		if err := config.DB.First(&user, &userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Attach user to the context
		c.Set("user", user)
		c.Next()
	}
}
