package controller

import (
	models "example/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// create a technology
func CreateTechnology(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var m models.Technology
	c.BindJSON(&m)
	db.Create(&m)
	c.JSON(200, m)
}
