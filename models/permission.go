package models

// only need mysql OR sqlite
// both are included here for reference
import (
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Name string `json:"name"`
}
