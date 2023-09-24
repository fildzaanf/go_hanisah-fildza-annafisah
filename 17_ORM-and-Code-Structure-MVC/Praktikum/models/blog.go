package models

import{
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
}

type Blog struct {
	gorm.Model
	UserId  int    `json:"userid" form:"userid"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}