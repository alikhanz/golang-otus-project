package resources

import (
	"context"

	"github.com/alikhanz/golang-otus-project/internal/config"
	"github.com/jinzhu/gorm"
)

type Resources struct {
	Db     *gorm.DB
	Config *config.Config
}

func Get(ctx context.Context) *Resources {
	r := &Resources{}
	r.initDb(ctx)

	return r
}
