package resources

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

func (r *Resources) initAMQP(ctx context.Context) {
	uri := amqp.URI{
		Scheme:   "amqp",
		Host:     r.Config.AMQPHost,
		Port:     r.Config.AMQPPort,
		Username: r.Config.AMQPLogin,
		Password: r.Config.AMQPPassword,
		Vhost:    r.Config.AMQPVhost,
	}

	conn, err := amqp.Dial(uri.String())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed connect to amqp")
	}

	r.AMQP = conn

	go func() {
		<-ctx.Done()
		log.Info().Msg("The db connection closing...")
		_ = conn.Close()
		log.Info().Msg("The db connection was successfully closed")
	}()
}
