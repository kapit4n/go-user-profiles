package controller

import (
	models "example/models"
	"fmt"

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

func UpdateRole(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var role models.Role
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&role)
	db.Save(&role)
	c.JSON(200, role)
}

func GetRole(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	var list []models.Role

	if err := db.Find(&list).Preload("Permissions").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, list)
	}
}

func GetRoleById(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	var role models.Role

	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).Preload("Permissions").First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.JSON(200, role)
}

// Assign role to person
func AssignPermission(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var role models.Role
	id := c.Params.ByName("id")

	// get role by id
	if err := db.Where("id = ?", id).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	db = db.Model(&role).Preload("Permissions")
	var permission models.Permission

	c.BindJSON(&permission)

	//get permission by id
	techId := permission.ID
	if err := db.Where("id = ?", techId).First(&permission).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	role.Permissions = append(role.Permissions, permission)

	db.Save(&role)
	c.JSON(200, role)
}
