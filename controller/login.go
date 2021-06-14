package controller

import (
	models "example/models"
	"example/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json providers")
		return
	}

	// find user with username
	err := models.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&u).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, "Please provide a valid login details")
		return
	}

	token := service.NewJWTService().GenerateToken(u.Username, true)

	c.JSON(http.StatusOK, token)
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, "Successfully logged out")
}
