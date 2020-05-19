package rotator

import (
	"context"

	"github.com/streadway/amqp"
	"golang.org/x/sync/errgroup"
)

type Config struct {
	AmqpHost     	 string
	AmqpPort     	 int
	AmqpLogin    	 string
	AmqpPassword 	 string
	AmqpVhost    	 string
	DbHost           string
	DbPort           int
	DbName           string
	DbUser           string
	DbPassword       string
	rmqConnection	 *amqp.Connection
}

type Rotator struct {
	Config Config
}

func NewRotator(c Config) *Rotator {
	return &Rotator{Config: c}
}

func (r *Rotator) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(r.connectRmq)

	return eg.Wait()
}

func (r *Rotator) connectRmq() error {
	uri := amqp.URI{
		Scheme:   "amqp",
		Host:     r.Config.AmqpHost,
		Port:     r.Config.AmqpPort,
		Username: r.Config.AmqpLogin,
		Password: r.Config.AmqpPassword,
		Vhost:    r.Config.AmqpVhost,
	}

	conn, err := amqp.Dial(uri.String())
	//failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	return err
}

//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Fatal().Msgf("%s: %s", msg, err)
//	}
//}