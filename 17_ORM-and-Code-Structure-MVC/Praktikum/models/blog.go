package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	UserId  uint   `json:"userid" form:"userid"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	User    User   `gorm:"foreignkey:UserId"`
}
