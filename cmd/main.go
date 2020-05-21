//go:generate ../scripts/protoc-generate.sh

package main

import (
	"context"

	"github.com/alikhanz/golang-otus-project/internal/helper"
	"github.com/alikhanz/golang-otus-project/internal/resources"
	"github.com/alikhanz/golang-otus-project/internal/rotator"
	"github.com/alikhanz/golang-otus-project/internal/server"
	"github.com/alikhanz/golang-otus-project/pkg/os_context"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := os_context.Context(context.Background())
	res := resources.Get(ctx)

	helper.Migrate(res)

	rr := createRotator(res)
	initServer(ctx, rr, res)
}


func initServer(ctx context.Context, r *rotator.Rotator, res *resources.Resources) {
	s := server.NewServer(
		r,
		res,
	)
	log.Fatal().Err(s.Run(ctx))
}

func createRotator(res *resources.Resources) *rotator.Rotator {
	return rotator.NewRotator(res)
}