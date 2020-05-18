package model

import "github.com/jinzhu/gorm"

type Banner struct {
	gorm.Model
	Id          int
	Description string
}