package models

// only need mysql OR sqlite
// both are included here for reference
import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Technology struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
