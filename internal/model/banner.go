package model

import "github.com/jinzhu/gorm"

type Banner struct {
	gorm.Model
	Description string
	Slots       []*Slot `gorm:"many2many:banner_slots;"`
}
