package controller

import (
	models "example/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// create a Role
func CreateRole(c *gin.Context) {
	var m models.Role
	c.BindJSON(&m)
	models.DB.Create(&m)
	c.JSON(200, m)
}

func UpdateRole(c *gin.Context) {
	var role models.Role
	id := c.Params.ByName("id")
	if err := models.DB.Where("id = ?", id).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&role)
	models.DB.Save(&role)
	c.JSON(200, role)
}

func GetRole(c *gin.Context) {
	var list []models.Role

	if err := models.DB.Find(&list).Preload("Permissions").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, list)
	}
}

func GetRoleById(c *gin.Context) {
	var role models.Role

	id := c.Params.ByName("id")

	if err := models.DB.Where("id = ?", id).Preload("Permissions").First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.JSON(200, role)
}

// Assign role to person
func AssignPermission(c *gin.Context) {
	var role models.Role
	id := c.Params.ByName("id")

	// get role by id
	if err := models.DB.Where("id = ?", id).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	db := models.DB.Model(&role).Preload("Permissions")
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

func DeleteRole(c *gin.Context) {
	id := c.Params.ByName("id")

	var toDelete models.Role
	models.DB.Where("id = ?", id).Delete(&toDelete)

	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
