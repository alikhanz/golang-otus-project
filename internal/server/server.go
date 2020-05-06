package server

import (
	"fmt"
	"github.com/alikhanz/golang-otus-project/pkg/pb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Config struct {
	GrpcPort int
}

type Server struct {
	Config        Config
	rotatorServer *RotatorServer
}

func NewServer(config Config) *Server {
	s := &Server{Config: config}
	s.rotatorServer = NewRotatorServer()

	return s
}

func (s *Server) Run() error {
	eg := errgroup.Group{}
	eg.Go(s.grpcInit)

	return eg.Wait()
}

func (s *Server) grpcInit() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRotatorServer(grpcServer, s.rotatorServer)
	return grpcServer.Serve(lis)
}

type RotatorServer struct {
	validator *Validator
}

func NewRotatorServer() *RotatorServer {
	v := NewValidator()
	return &RotatorServer{validator: v}
}
