package controller

import (
	models "example/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person models.Person
	models.DB.Where("id = ?", id).Delete(&person)

	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// update a person
func UpdatePerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")
	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	models.DB.Save(&person)
	c.JSON(200, person)
}

// Assign role to person
func AssignRole(c *gin.Context) {
	var role models.Role
	var person models.Person
	id := c.Params.ByName("id")
	// get person by id
	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	models.DB.Model(&person).Preload("Roles")

	c.BindJSON(&role)

	//get role by id
	roleId := role.ID
	if err := models.DB.Where("id = ?", roleId).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	person.Roles = append(person.Roles, role)

	models.DB.Save(&person)
	c.JSON(200, person)
}

// Assign role to person
func AssignExperience(c *gin.Context) {
	var experience models.Experience
	var person models.Person
	id := c.Params.ByName("id")
	// get person by id
	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	models.DB.Model(&person).Preload("Experiences")

	c.BindJSON(&experience)

	//get experience by id
	experienceId := experience.ID
	if err := models.DB.Where("id = ?", experienceId).First(&experience).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	person.Experiences = append(person.Experiences, experience)

	models.DB.Save(&person)
	c.JSON(200, person)
}

// UnAssign role to person
func UnAssignRole(c *gin.Context) {
	var role models.Role
	var person models.Person
	id := c.Params.ByName("id")
	// get person by id
	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&role)

	//get role by id
	roleId := role.ID
	if err := models.DB.Where("id = ?", roleId).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	models.DB.Model(&person).Association("Roles").Delete(&role)

	c.JSON(200, person)
}

// create a person
func CreatePerson(c *gin.Context) {
	var person models.Person
	c.BindJSON(&person)

	var role models.Role
	roleId := person.RoleId
	if err := models.DB.Where("id = ?", roleId).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	var roles []models.Role

	roles = append(roles, role)

	person.Roles = roles

	models.DB.Create(&person)
	c.JSON(200, person)
}

func GetPersonById(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")

	if err := models.DB.First(&person, id).Preload("Roles").Preload("Experiences").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

func GetPersons(c *gin.Context) {
	var list []models.Person
	if err := models.DB.Find(&list).Preload("Roles").Preload("Experiences").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, list)
	}
}
