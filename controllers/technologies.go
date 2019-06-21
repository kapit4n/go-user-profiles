package controllers

import (
	models "example/models"

	"github.com/gin-gonic/gin"
)

// create a technology
func CreateTechnology(c *gin.Context) {
	var m models.Technology
	c.BindJSON(&m)
	db.Create(&m)
	c.JSON(200, m)
}
