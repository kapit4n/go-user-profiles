package controller

import (
	models "example/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// create a Permission
func CreatePermissions(c *gin.Context) {
	var m models.Permission
	c.BindJSON(&m)
	models.DB.Create(&m)
	c.JSON(200, m)
}

func UpdatePermissions(c *gin.Context) {
	var permission models.Permission
	id := c.Params.ByName("id")
	if err := models.DB.Where("id = ?", id).First(&permission).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&permission)
	models.DB.Save(&permission)
	c.JSON(200, permission)
}

func GetPermissions(c *gin.Context) {
	var list []models.Permission

	if err := models.DB.Find(&list).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, list)
	}
}

func GetPermissionById(c *gin.Context) {
	var permission models.Permission
	id := c.Params.ByName("id")

	if err := models.DB.First(&permission, id).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, permission)
	}
}

func DeletePermission(c *gin.Context) {
	id := c.Params.ByName("id")

	var toDelete models.Permission
	models.DB.Where("id = ?", id).Delete(&toDelete)

	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
