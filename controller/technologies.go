package controller

import (
	models "example/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateTechnology(c *gin.Context) {
	var m models.Technology
	c.BindJSON(&m)
	models.DB.Create(&m)
	c.JSON(200, m)
}

func GetTechnology(c *gin.Context) {
	var technologies []models.Technology

	if err := models.DB.Find(&technologies).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, technologies)
	}
}

func GetTechnologyById(c *gin.Context) {
	var technology []models.Technology
	id := c.Params.ByName("id")

	if err := models.DB.First(&technology, id).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, technology)
	}
}

func DeleteTechnology(c *gin.Context) {
	id := c.Params.ByName("id")

	var toDelete models.Technology
	models.DB.Where("id = ?", id).Delete(&toDelete)

	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
