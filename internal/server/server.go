package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/alikhanz/golang-otus-project/internal/resources"
	"github.com/alikhanz/golang-otus-project/internal/rotator"
	"github.com/alikhanz/golang-otus-project/pkg/pb"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Server struct {
	Resources     *resources.Resources
	rotatorServer *rotator.Server
}

func NewServer(r *rotator.Rotator, res *resources.Resources) *Server {
	s := &Server{
		Resources: res,
	}
	s.rotatorServer = rotator.NewRotatorServer(r)

	return s
}

func (s *Server) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return s.grpcInit(ctx)
	})

	return eg.Wait()
}

func (s *Server) grpcInit(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Resources.Config.GrpcPort))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRotatorServer(grpcServer, s.rotatorServer)

	go func() {
		go func() {
			<-ctx.Done()
			log.Info().Msg("The grpc-server is shutting down...")
			grpcServer.GracefulStop()
			log.Info().Msg("The grpc-server was successfully stopped")
		}()
		<-ctx.Done()
		time.Sleep(time.Second * 10)
		grpcServer.Stop()
		log.Info().Msg("The grpc-server was successfully stopped")
	}()

	return grpcServer.Serve(lis)
}
