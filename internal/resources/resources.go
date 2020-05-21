package resources

import (
	"context"

	"github.com/alikhanz/golang-otus-project/internal/config"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

type Resources struct {
	DB     *gorm.DB
	Config *config.Config
	AMQP   *amqp.Connection
}

func Get(ctx context.Context) *Resources {
	r := &Resources{}
	r.initConfig()
	r.initDB(ctx)
	r.initAMQP(ctx)

	return r
}
