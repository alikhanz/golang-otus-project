package resources

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

func (r *Resources) initAmqp(ctx context.Context) {
	uri := amqp.URI{
		Scheme:   "amqp",
		Host:     r.Config.AmqpHost,
		Port:     r.Config.AmqpPort,
		Username: r.Config.AmqpLogin,
		Password: r.Config.AmqpPassword,
		Vhost:    r.Config.AmqpVhost,
	}

	conn, err := amqp.Dial(uri.String())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed connect to amqp")
	}

	r.Amqp = conn

	go func() {
		<-ctx.Done()
		log.Info().Msg("The db connection closing...")
		_ = conn.Close()
		log.Info().Msg("The db connection was successfully closed")
	}()
}