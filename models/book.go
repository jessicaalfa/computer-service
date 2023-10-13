package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID       int    `json:"id" form:"id"`
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}
