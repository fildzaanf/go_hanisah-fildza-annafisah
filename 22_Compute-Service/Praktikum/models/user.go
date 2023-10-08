package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Blogs    []Blog `gorm:"ForeignKey:UserId;references:ID"`
}
