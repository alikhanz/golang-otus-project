package model

import "time"

type Stat struct {
	ShowCount int64 `gorm:"DEFAULT:0"`
	HitCount  int64 `gorm:"DEFAULT:0"`
	BannerID  uint  `gorm:"primary_key;auto_increment:false"`
	SlotID    uint  `gorm:"primary_key;auto_increment:false"`
	SdgID     uint  `gorm:"primary_key;auto_increment:false"`
	Banner    Banner
	Slot      Slot
	Sdg       Sdg
	CreatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
