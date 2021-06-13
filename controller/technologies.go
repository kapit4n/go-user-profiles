package controller

import (
	models "example/models"
	"fmt"

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

func GetTechnology(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var technologies []models.Technology

	if err := db.Find(&technologies).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, technologies)
	}
}

func GetTechnologyById(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var technology []models.Technology
	id := c.Params.ByName("id")

	if err := db.First(&technology, id).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, technology)
	}
}

func DeleteTechnology(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	id := c.Params.ByName("id")

	var toDelete models.Technology
	db.Where("id = ?", id).Delete(&toDelete)

	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
