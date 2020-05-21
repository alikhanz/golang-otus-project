package resources

import (
	"context"

	"github.com/alikhanz/golang-otus-project/internal/config"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

type Resources struct {
	Db     *gorm.DB
	Config *config.Config
	Amqp   *amqp.Connection
}

func Get(ctx context.Context) *Resources {
	r := &Resources{}
	r.initConfig()
	r.initDb(ctx)

	return r
}
