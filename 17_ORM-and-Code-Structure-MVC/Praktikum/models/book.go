package models

import{
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
}

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}