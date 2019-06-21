package main

// only need mysql OR sqlite
// both are included here for reference
import (
	models "example/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	//db, err = gorm.Open("sqlite3", "./gorm.db")
	db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Experience{})
	db.AutoMigrate(&models.Technology{})
	r := gin.Default()
	r.GET("/people/", GetPeople)
	r.GET("/people/:id", GetPerson)
	r.POST("/people", CreatePerson)
	r.POST("/technology", CreateTechnology)
	r.POST("/people/:id/experience", CreateExperience)
	r.POST("/experience/:id/atech", AssignTech)
	r.POST("/people/:id/arole", AssignRole)
	r.POST("/people/:id/aexperience", AssignExperience)
	r.POST("/people/:id/urole", UnAssignRole)
	r.PUT("/people/:id", UpdatePerson)
	r.DELETE("/people/:id", DeletePerson)
	r.Run(":8080")
}

// delete a person
func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person models.Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// update a person
func UpdatePerson(c *gin.Context) {
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

// Assign role to person
func AssignTech(c *gin.Context) {
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

// UnAssign role to person
func UnAssignRole(c *gin.Context) {
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

// create a technology
func CreateTechnology(c *gin.Context) {
	var m models.Technology
	c.BindJSON(&m)
	db.Create(&m)
	c.JSON(200, m)
}

// get a person by id
func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person models.Person
	db = db.Model(&person).Preload("Roles").Preload("Experiences")
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

// get all people
func GetPeople(c *gin.Context) {
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
