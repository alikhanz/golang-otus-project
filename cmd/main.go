//go:generate ../scripts/protoc-generate.sh

package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alikhanz/golang-otus-project/internal/resources"
	rotator2 "github.com/alikhanz/golang-otus-project/internal/rotator"
	"github.com/alikhanz/golang-otus-project/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := osContext(context.Background())
	res := resources.Get(ctx)

	rr := createRotator(res)
	initServer(ctx, rr, res)
}

var (
	defaultSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
)

func osContext(ctx context.Context, signals ...os.Signal) context.Context {
	if len(signals) == 0 {
		signals = defaultSignals
	}

	ctx, cancel := context.WithCancel(ctx)
	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, signals...)
	go func() {
		select {
		case <-ctx.Done():
		case sig := <-signalsChan:
			log.Info().Msgf("Got signal %s", sig)
			cancel()
		}
	}()
	return ctx
}

func initServer(ctx context.Context, r *rotator2.Rotator, res *resources.Resources) {
	s := server.NewServer(
		r,
		res,
	)
	log.Fatal().Err(s.Run(ctx))
}

func createRotator(res *resources.Resources) *rotator2.Rotator {
	return rotator2.NewRotator(rotator2.Config{
		AmqpHost:     res.Config.AmqpHost,
		AmqpPort:     res.Config.AmqpPort,
		AmqpLogin:    res.Config.AmqpLogin,
		AmqpPassword: res.Config.AmqpPassword,
		AmqpVhost:    res.Config.AmqpVhost,
		DbHost:       res.Config.DbHost,
		DbPort:       res.Config.DbPort,
		DbName:       res.Config.DbName,
		DbUser:       res.Config.DbUser,
		DbPassword:   res.Config.DbPassword,
	})
}