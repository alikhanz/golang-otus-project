package resources

import (
	"context"

	"github.com/alikhanz/golang-otus-project/internal/config"
	"github.com/jinzhu/gorm"
)

type Resources struct {
	DB     *gorm.DB
	Config *config.Config
}

func Get(ctx context.Context) *Resources {
	r := &Resources{}
	r.initConfig()
	r.initDB(ctx)

	return r
}
