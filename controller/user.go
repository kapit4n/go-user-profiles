package controller

import (
	models "example/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Input valid json")
	}

	models.DB.Create(&u)
	c.JSON(200, &u)
}

func GetUsers(c *gin.Context) {
	var list []models.User

	if err := models.DB.Find(&list).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, list)
	}

}
