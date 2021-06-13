package controller

import (
	models "example/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Assign role to person
func AssignTech(c *gin.Context) {
	var experience models.Experience
	id := c.Params.ByName("id")
	// get experience by id
	if err := models.DB.Where("id = ?", id).First(&experience).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	models.DB.Model(&experience).Preload("Technologies")
	var technology models.Technology

	c.BindJSON(&technology)

	//get technology by id
	techId := technology.ID
	if err := models.DB.Where("id = ?", techId).First(&technology).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	experience.Technologies = append(experience.Technologies, technology)

	models.DB.Save(&experience)
	c.JSON(200, experience)
}

// Assign role to person
func CreateExperience(c *gin.Context) {
	var experience models.Experience
	var person models.Person
	c.BindJSON(&experience)

	// get person by id
	id := c.Params.ByName("id")
	if experience.PersonID > 0 {
		id = strconv.FormatUint(uint64(experience.PersonID), 10)
	}

	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	experience.Person = person
	models.DB.Create(&experience)
	c.JSON(200, experience)
}

func UpdateExperience(c *gin.Context) {
	var experience models.Experience
	id := c.Params.ByName("id")
	if err := models.DB.Where("id = ?", id).First(&experience).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&experience)
	models.DB.Save(&experience)
	c.JSON(200, experience)
}

func GetExperience(c *gin.Context) {
	var list []models.Experience

	if err := models.DB.Find(&list).Preload("Person").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, list)
	}
}

func GetExperienceById(c *gin.Context) {
	var experience models.Experience
	id := c.Params.ByName("id")

	if err := models.DB.Where("id = ?", id).Preload("Technologies").First(&experience).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, experience)
}

func DeleteExperience(c *gin.Context) {
	id := c.Params.ByName("id")

	var toDelete models.Experience
	models.DB.Where("id = ?", id).Delete(&toDelete)

	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
