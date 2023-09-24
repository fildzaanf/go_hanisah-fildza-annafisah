package models

import{
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Blogs    []Blog `gorm:"ForeignKey:UserId;references:id"`
}
  