package helper

import (
	"github.com/alikhanz/golang-otus-project/internal/model"
	"github.com/alikhanz/golang-otus-project/internal/resources"
)

func Migrate(res *resources.Resources) {
	res.Db.AutoMigrate(&model.Banner{}, &model.Sdg{}, &model.Slot{}, &model.Stat{})
}