package gormginintro

import (
  	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name		string	`json:"name"`
	Age			int		`json:"age"`
}

type Group struct {
	gorm.Model
	Name	string	`json:"name"`
}