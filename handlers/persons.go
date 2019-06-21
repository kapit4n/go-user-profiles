package handlers

import (
	models "example/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// delete a person
func DeletePerson(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	id := c.Params.ByName("id")
	var person models.Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// update a person
func UpdatePerson(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	var person models.Person
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	db.Save(&person)
	c.JSON(200, person)
}

// Assign role to person
func CreateExperience(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	var experience models.Experience
	var person models.Person
	id := c.Params.ByName("id")
	// get person by id
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&experience)
	experience.Person = person
	db.Create(&experience)
	c.JSON(200, experience)
}

// Assign role to person
func AssignRole(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	var role models.Role
	var person models.Person
	id := c.Params.ByName("id")
	// get person by id
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	db = db.Model(&person).Preload("Roles")

	c.BindJSON(&role)

	//get role by id
	roleId := role.ID
	if err := db.Where("id = ?", roleId).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	person.Roles = append(person.Roles, role)

	db.Save(&person)
	c.JSON(200, person)
}

// Assign role to person
func AssignExperience(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	var experience models.Experience
	var person models.Person
	id := c.Params.ByName("id")
	// get person by id
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	db = db.Model(&person).Preload("Experiences")

	c.BindJSON(&experience)

	//get experience by id
	experienceId := experience.ID
	if err := db.Where("id = ?", experienceId).First(&experience).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	person.Experiences = append(person.Experiences, experience)

	db.Save(&person)
	c.JSON(200, person)
}

// UnAssign role to person
func UnAssignRole(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	var role models.Role
	var person models.Person
	id := c.Params.ByName("id")
	// get person by id
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&role)

	//get role by id
	roleId := role.ID
	if err := db.Where("id = ?", roleId).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	db.Model(&person).Association("Roles").Delete(&role)

	c.JSON(200, person)
}

// create a person
func CreatePerson(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	var person models.Person
	c.BindJSON(&person)

	log.Println(person)

	var role models.Role
	roleId := person.RoleId
	if err := db.Where("id = ?", roleId).First(&role).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	var roles []models.Role

	roles = append(roles, role)

	person.Roles = roles

	db.Create(&person)
	c.JSON(200, person)
}

// get a person by id
func GetPerson(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")

	id := c.Params.ByName("id")
	person := new(models.Person)
	db = db.Model(person).Preload("Roles").Preload("Experiences")
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

// get all people
func GetPeople(c *gin.Context) {
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	var people []models.Person
	roles := []models.Role{}

	db = db.Model(&people).Preload("Roles")
	db = db.Find(&people)

	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {

		log.Println(people)

		log.Println(roles)

		c.JSON(200, people)
	}
}
