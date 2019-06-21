package handlers

import (
	models "example/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// create a technology
func CreateTechnology(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	var m models.Technology
	c.BindJSON(&m)
	db.Create(&m)
	c.JSON(200, m)
}
