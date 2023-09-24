package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}
