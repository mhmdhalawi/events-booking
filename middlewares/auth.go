package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhmdhalawi/events-booking/utils"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	userID, err := utils.ValidateToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	c.Set("userID", userID)

	c.Next()
}
