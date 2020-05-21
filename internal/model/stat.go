package model

import "time"

type Stat struct {
	ShowCount int64 `gorm:"DEFAULT:0"`
	HitCount  int64 `gorm:"DEFAULT:0"`
	BannerId  uint  `gorm:"primary_key;auto_increment:false"`
	SlotId    uint  `gorm:"primary_key;auto_increment:false"`
	SdgId     uint  `gorm:"primary_key;auto_increment:false"`
	Banner    Banner
	Slot      Slot
	Sdg       Sdg
	CreatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
