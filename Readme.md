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

## TODO
* CRUD persons
* CRUD technologies
* CRUD experiences
* CRUD roles

## TOOLS
### Add people
* curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"firstname":"Luis","lastname":"Arce", "city": "Cochabamba"}' \
  http://localhost:8080/people

### Add technologies
* curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Golang"}' \
  http://localhost:8080/technologies

* curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"React js"}' \
  http://localhost:8080/technologies

### Add roles
* curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Admin"}' \
  http://localhost:8080/roles

* curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Developer"}' \
  http://localhost:8080/roles

* curl --header "Content-Type: application/json" \
  --request PUT \
  --data '{"name":"Super Admin"}' \
  http://localhost:8080/roles/1


### Add permissions
* curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Create User"}' \
  http://localhost:8080/permissions

* curl --header "Content-Type: application/json" \
  --request PUT \
  --data '{"name":"Create User Permission"}' \
  http://localhost:8080/permissions/1


### Experiences
* curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"title":"Remote Software developer", "personId": 1}' \
  http://localhost:8080/experiences

* curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"id": 2}' \
  http://localhost:8080/experiences/1/atech




## before add swagger
- go get -u github.com/swaggo/swag/cmd/swag
- export PATH=$(go env GOPATH)/bin:$PATH
- https://github.com/swaggo/gin-swagger
