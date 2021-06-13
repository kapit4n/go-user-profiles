package controller

import (
	models "example/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Assign role to person
func AssignTech(c *gin.Context) {
	// db, _ = gorm.Open("mysql", "root:root@tcp(localhost:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var experience models.Experience
	id := c.Params.ByName("id")
	// get experience by id
	if err := db.Where("id = ?", id).First(&experience).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	db = db.Model(&experience).Preload("Technologies")
	var technology models.Technology

	c.BindJSON(&technology)

	//get technology by id
	techId := technology.ID
	if err := db.Where("id = ?", techId).First(&technology).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	experience.Technologies = append(experience.Technologies, technology)

	db.Save(&experience)
	c.JSON(200, experience)
}

// Assign role to person
func CreateExperience(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var experience models.Experience
	var person models.Person
	c.BindJSON(&experience)

	// get person by id
	id := c.Params.ByName("id")
	if experience.PersonID > 0 {
		id = strconv.FormatUint(uint64(experience.PersonID), 10)
	}

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	experience.Person = person
	db.Create(&experience)
	c.JSON(200, experience)
}

func UpdateExperience(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	var experience models.Experience
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&experience).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&experience)
	db.Save(&experience)
	c.JSON(200, experience)
}

func GetExperience(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	var list []models.Experience

	if err := db.Find(&list).Preload("Person").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, list)
	}
}

func GetExperienceById(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	var experience models.Experience
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).Preload("Technologies").First(&experience).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, experience)
}
