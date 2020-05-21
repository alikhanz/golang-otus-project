package model

import "github.com/jinzhu/gorm"

type Slot struct {
	gorm.Model
	Description string
	Banners     []*Banner `gorm:"many2many:banner_slots;"`
}
