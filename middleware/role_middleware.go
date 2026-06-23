package middleware

import (
	"Isong/config"
	"Isong/models"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "user not found"})
			return
		}

		if user.Role != requiredRole {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden: insufficient permissions"})
			return
		}

		c.Next()
	}
}
