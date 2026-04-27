package grpc

import (
	"fmt"
	"log"
	"net"

	userpb "github.com/Yusufdot101/ripple-proto/golang/user/v4"
	"github.com/Yusufdot101/ripple/services/user/config"
	"github.com/Yusufdot101/ripple/services/user/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	userpb.UnimplementedUserServiceServer
	port int
	asvc ports.AuthService
	usvc ports.UserService
}

func NewAdapter(port int, asvc ports.AuthService, usvc ports.UserService) *Adapter {
	return &Adapter{
		port: port,
		asvc: asvc,
		usvc: usvc,
	}
}

func (a *Adapter) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	log.Printf("grpc server listening on port :%d\n", a.port)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to server grpc server: %v", err)
	}

	return nil
}
