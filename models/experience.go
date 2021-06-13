package models

// only need mysql OR sqlite
// both are included here for reference
import (
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Title        string `json:"title"`
	Person       Person
	PersonID     uint         `json:"personId"`
	Technologies []Technology `gorm:"many2many:experience_tech"`
}
