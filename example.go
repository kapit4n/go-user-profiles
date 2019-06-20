package main
/*
import (
 "fmt"
 "github.com/gin-gonic/gin"
 "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Person struct {
 gorm.Model
 FirstName string `json:"firstname"`
 LastName string `json:"lastname"`
 RoleID uint `json:"roleId"`
 Role Role `gorm:"many2many:person_role;"`
}

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

func main() {
 // NOTE: See we’re using = to assign the global var
 // instead of := which would assign it only in this function
 db, err = gorm.Open("sqlite3", "./gorm.db")
 if err != nil {
    fmt.Println(err)
 }
defer db.Close()
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Role{})

	r := gin.Default()
	r.GET("/role/", GetRole)
	r.POST("/role", CreateRole)
	r.GET("/people/", GetPeople)
	r.POST("/people", CreatePerson)
	r.GET("/people/:id", GetPerson)
	r.PUT("/people/:id", UpdatePerson)

	r.Run(":8080")
}

func UpdatePerson(c *gin.Context) {
 var person Person
 id := c.Params.ByName("id")
 if err := db.Where("id = ?", id).First(&person).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 }
 c.BindJSON(&person)
 db.Save(&person)
 c.JSON(200, person)
}

func CreateRole(c *gin.Context) {
 var role Role
 c.BindJSON(&role)
 db.Create(&role)
 c.JSON(200, role)
}

func GetRole(c *gin.Context) {
	var roles []Role
	if err := db.Find(&roles).Error; err != nil {
	   c.AbortWithStatus(404)
	   fmt.Println(err)
	} else {
	   c.JSON(200, roles)
	}
}

func CreatePerson(c *gin.Context) {
 var person Person
 c.BindJSON(&person)
 db.Create(&person)
 c.JSON(200, person)
}

func GetPerson(c *gin.Context) {
 id := c.Params.ByName("id")
 var person Person
 if err := db.Where("id = ?", id).First(&person).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, person)
 }
}

func GetPeople(c *gin.Context) {
 var people []Person

 var users []Person
role := Role{}

db.First(&role, "id = ?", 111)

db.Model(&role).Related(&users,  "Users")

 if err := db.Find(&people).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, people)
 }
}

*/