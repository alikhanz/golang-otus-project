package model

import "github.com/jinzhu/gorm"

// Социально-демографическая группа

type Sdg struct {
	gorm.Model
	Id          int
	Description string
}
