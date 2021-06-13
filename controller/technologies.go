package controller

import (
	models "example/models"
	"fmt"
	"log"

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

	db = db.Find(&technologies)

	if err := db.Find(&technologies).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {

		log.Println(technologies)

		c.JSON(200, technologies)
	}

}
