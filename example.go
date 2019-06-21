package main

// only need mysql OR sqlite
// both are included here for reference
import (
	handlers "example/handlers"
	models "example/models"
	"fmt"

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
	r.GET("/people/", handlers.GetPeople)
	r.GET("/people/:id", handlers.GetPerson)
	r.POST("/people", handlers.CreatePerson)
	r.POST("/technology", handlers.CreateTechnology)
	r.POST("/people/:id/experience", handlers.CreateExperience)
	r.POST("/experience/:id/atech", handlers.AssignTech)
	r.POST("/people/:id/arole", handlers.AssignRole)
	r.POST("/people/:id/aexperience", handlers.AssignExperience)
	r.POST("/people/:id/urole", handlers.UnAssignRole)
	r.PUT("/people/:id", handlers.UpdatePerson)
	r.DELETE("/people/:id", handlers.DeletePerson)
	r.Run(":8080")
}
