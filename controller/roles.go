package controller

import (
	models "example/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// create a Role
func CreateRole(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var m models.Role
	c.BindJSON(&m)
	db.Create(&m)
	c.JSON(200, m)
}

func GetRole(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	var list []models.Role

	db = db.Find(&list)

	if err := db.Find(&list).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {

		log.Println(list)

		c.JSON(200, list)
	}

}
