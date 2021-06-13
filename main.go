package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"example/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	/* db, err = gorm.Open("sqlite3", "./gorm.db")
	//db, _ = gorm.Open("mysql", "root:root@tcp(localhost:3306)/peco")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Experience{})
	db.AutoMigrate(&models.Technology{})

	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	*/

	r := gin.New()
	r.GET("/people", controller.GetPeople)
	r.GET("/people/:id", controller.GetPerson)
	r.POST("/people", controller.CreatePerson)
	r.POST("/people/:id/experience", controller.CreateExperience)
	r.POST("/people/:id/arole", controller.AssignRole)
	r.POST("/people/:id/aexperience", controller.AssignExperience)
	r.POST("/people/:id/urole", controller.UnAssignRole)
	r.PUT("/people/:id", controller.UpdatePerson)
	r.DELETE("/people/:id", controller.DeletePerson)

	r.GET("/technologies", controller.GetTechnology)
	r.POST("/technologies", controller.CreateTechnology)

	r.GET("/roles", controller.GetRole)
	r.POST("/roles", controller.CreateRole)
	r.PUT("/roles/:id", controller.UpdateRole)

	r.GET("/experiences", controller.GetExperience)
	r.POST("/experiences", controller.CreateExperience)
	r.PUT("/experiences/:id", controller.UpdateExperience)

	r.POST("/experience/:id/atech", controller.AssignTech)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()

	//r.Run(":8080")
}
