package model

import "github.com/jinzhu/gorm"

type Stat struct {
	gorm.Model
	ShowCount int64
	HitCount  int64
	BannerId  int
	SlotId    int
	SdgId	  int
}