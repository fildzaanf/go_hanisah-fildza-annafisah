package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}
