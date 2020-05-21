package model

import "github.com/jinzhu/gorm"

// Социально-демографическая группа
type Sdg struct {
	gorm.Model
	Description string
}