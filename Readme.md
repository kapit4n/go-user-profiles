# Go example with GORM and GIN

## PREREQUISITES
* Have a db created with peco name
* have a user with permission to create tables and alter them
* go get "github.com/gin-gonic/gin"
* go get "github.com/go-sql-driver/mysql" 
* go get "github.com/jinzhu/gorm"
* go get "github.com/jinzhu/gorm/dialects/sqlite"

## RUN
go run example.go

## RESOURCES
```
[GIN-debug] GET    /people/                  --> main.GetPeople (3 handlers)
[GIN-debug] GET    /people/:id               --> main.GetPerson (3 handlers)
[GIN-debug] POST   /people                   --> main.CreatePerson (3 handlers)
[GIN-debug] POST   /technology               --> main.CreateTechnology (3 handlers)
[GIN-debug] POST   /people/:id/experience    --> main.CreateExperience (3 handlers)
[GIN-debug] POST   /experience/:id/atech     --> main.AssignTech (3 handlers)
[GIN-debug] POST   /people/:id/arole         --> main.AssignRole (3 handlers)
[GIN-debug] POST   /people/:id/aexperience   --> main.AssignExperience (3 handlers)
[GIN-debug] POST   /people/:id/urole         --> main.UnAssignRole (3 handlers)
[GIN-debug] PUT    /people/:id               --> main.UpdatePerson (3 handlers)
[GIN-debug] DELETE /people/:id               --> main.DeletePerson (3 handlers)
```

## Structure
* /models
* /handlers
