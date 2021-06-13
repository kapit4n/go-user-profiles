package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"example/controller"
	"example/models"

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
	//db, _ = gorm.Open("mysql", "root:root@tcp(localhost:3306)/peco")

	models.ConnectionDB()

	r := gin.New()
	r.GET("/people", controller.GetPersons)
	r.GET("/people/:id", controller.GetPersonById)
	r.POST("/people", controller.CreatePerson)
	r.PUT("/people/:id", controller.UpdatePerson)
	r.DELETE("/people/:id", controller.DeletePerson)

	r.POST("/people/:id/experience", controller.CreateExperience)
	r.POST("/people/:id/arole", controller.AssignRole)
	r.POST("/people/:id/aexperience", controller.AssignExperience)
	r.POST("/people/:id/urole", controller.UnAssignRole)

	r.GET("/technologies", controller.GetTechnology)
	r.GET("/technologies/:id", controller.GetTechnologyById)
	r.POST("/technologies", controller.CreateTechnology)
	r.DELETE("/technologies/:id", controller.DeleteTechnology)

	r.GET("/roles", controller.GetRole)
	r.GET("/roles/:id", controller.GetRoleById)
	r.POST("/roles", controller.CreateRole)
	r.POST("/roles/:id/assignPermission", controller.AssignPermission)
	r.PUT("/roles/:id", controller.UpdateRole)
	r.DELETE("/roles/:id", controller.DeleteRole)

	r.GET("/permissions", controller.GetPermissions)
	r.GET("/permissions/:id", controller.GetPermissionById)
	r.POST("/permissions", controller.CreatePermissions)
	r.PUT("/permissions/:id", controller.UpdatePermissions)
	r.DELETE("/permissions/:id", controller.DeletePermission)

	r.GET("/experiences", controller.GetExperience)
	r.GET("/experiences/:id", controller.GetExperienceById)
	r.POST("/experiences", controller.CreateExperience)
	r.PUT("/experiences/:id", controller.UpdateExperience)
	r.POST("/experiences/:id/atech", controller.AssignTech)
	r.DELETE("/experiences/:id", controller.DeleteExperience)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()
}
