package models

// only need mysql OR sqlite
// both are included here for reference
import (
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
	ID          uint         `json:"id"`
	FirstName   string       `json:"firstname"`
	LastName    string       `json:"lastname"`
	City        string       `json:"city"`
	Roles       []Role       `gorm:"many2many:person_roles;"`
	RoleId      uint         `json:"roleId"`
	Experiences []Experience `gorm:"many2many:person_exp;PRELOAD:false;"`
}
