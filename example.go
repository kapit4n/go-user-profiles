package main
// only need mysql OR sqlite 
// both are included here for reference
import (
 "fmt"
 "log"
 "github.com/gin-gonic/gin"
 _ "github.com/go-sql-driver/mysql" 
 "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/sqlite"
)
var db *gorm.DB
var err error
type Person struct {
 ID uint `json:"id"`
 FirstName string `json:"firstname"`
 LastName string `json:"lastname"`
 City string `json:"city"`
 Roles []Role `gorm:"many2many:person_roles;"`
 RoleId uint `json:"roleId"`
}

type Role struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

func main() {
 // NOTE: See we’re using = to assign the global var
 // instead of := which would assign it only in this function
 //db, err = gorm.Open("sqlite3", "./gorm.db")
 db, _ = gorm.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/peco?charset=utf8&parseTime=True&loc=Local")
 if err != nil {
    fmt.Println(err)
 }
 defer db.Close()
 db.AutoMigrate(&Person{})
 db.AutoMigrate(&Role{})
 r := gin.Default()
 r.GET("/people/", GetPeople)
 r.GET("/people/:id", GetPerson)
 r.POST("/people", CreatePerson)
 r.POST("/people/:id/role", AssignRole)
 r.PUT("/people/:id", UpdatePerson)
 r.DELETE("/people/:id", DeletePerson)
 r.Run(":8080")
}

// delete a person
func DeletePerson(c *gin.Context) {
 id := c.Params.ByName("id")
 var person Person
 d := db.Where("id = ?", id).Delete(&person)
 fmt.Println(d)
 c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// update a person
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

// Assign role to person
func AssignRole(c *gin.Context) {
 var role Role
 var person Person
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
 var roles []Role

 roles = append(roles, role)
 person.Roles = roles

 db.Save(&person)
 c.JSON(200, person)
}

// create a person
func CreatePerson(c *gin.Context) {
 var person Person
 c.BindJSON(&person)
 
 log.Println(person)
 
 var role Role
 roleId := person.RoleId
 if err := db.Where("id = ?", roleId).First(&role).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 }
 var roles []Role

 roles = append(roles, role)

 person.Roles = roles

 db.Create(&person)
 c.JSON(200, person)
}

// get a person by id
func GetPerson(c *gin.Context) {
 id := c.Params.ByName("id")
 var person Person
 db = db.Model(&person).Preload("Roles")
 if err := db.Where("id = ?", id).First(&person).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, person)
 }
}

// get all people
func GetPeople(c *gin.Context) {
 var people []Person
 roles := []Role{}
 
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