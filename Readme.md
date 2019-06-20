# Go example with GORM and GIN

## RUN
go run example2.go

## RESOURCES
```
[GIN-debug] GET    /people/                  --> main.GetPeople (3 handlers)
[GIN-debug] GET    /people/:id               --> main.GetPerson (3 handlers)
[GIN-debug] POST   /people                   --> main.CreatePerson (3 handlers)
[GIN-debug] POST   /people/:id/role          --> main.AssignRole (3 handlers)
[GIN-debug] PUT    /people/:id               --> main.UpdatePerson (3 handlers)
[GIN-debug] DELETE /people/:id               --> main.DeletePerson (3 handlers
```