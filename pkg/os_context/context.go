package os_context

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

var (
	defaultSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
)

func Context(ctx context.Context, signals ...os.Signal) context.Context {
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