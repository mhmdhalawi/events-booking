package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhmdhalawi/events-booking/models"
	"github.com/mhmdhalawi/events-booking/utils"
)

func signup(c *gin.Context) {

	var user models.User

	if c.ShouldBindJSON(&user) == nil {
		err := user.Save()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create user"})
			return
		}

		token, err := utils.GenerateToken(user.Email, user.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "user created", "token": token})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create user"})
	}

}

func login(c *gin.Context) {
	var user models.User

	if c.ShouldBindJSON(&user) == nil {
		userDB, err := user.FindByEmail()

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		if !utils.ComparePasswords(userDB.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
			return
		}

		token, err := utils.GenerateToken(userDB.Email, userDB.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot login"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "logged in", "token": token})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot login"})
	}

}

func authRoutes(route *gin.RouterGroup) {
	route.POST("/signup", signup)
	route.POST("/login", login)
}
