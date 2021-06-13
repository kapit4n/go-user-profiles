package controller

import (
	models "example/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// create a Permission
func CreatePermissions(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var m models.Permission
	c.BindJSON(&m)
	db.Create(&m)
	c.JSON(200, m)
}

func UpdatePermissions(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var permission models.Permission
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&permission).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&permission)
	db.Save(&permission)
	c.JSON(200, permission)
}

func GetPermissions(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	var list []models.Permission

	if err := db.Find(&list).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, list)
	}
}

func GetPermissionById(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	var permission models.Permission
	id := c.Params.ByName("id")

	if err := db.First(&permission, id).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, permission)
	}
}
