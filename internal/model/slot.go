package model

import "github.com/jinzhu/gorm"

type Slot struct {
	gorm.Model
	Id          int
	Description string
}