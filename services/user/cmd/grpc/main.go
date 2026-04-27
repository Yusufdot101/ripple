package main

import (
	"context"
	"log"
	"time"

	"github.com/Yusufdot101/ripple/services/user/config"
	"github.com/Yusufdot101/ripple/services/user/internal/adapters/primary/grpc"
	"github.com/Yusufdot101/ripple/services/user/internal/adapters/secondary/google"
	"github.com/Yusufdot101/ripple/services/user/internal/adapters/secondary/postgresql"
	"github.com/Yusufdot101/ripple/services/user/internal/application/core/services"
)

func main() {
	// get repo
	repo, err := postgresql.NewAdapter(config.GetDatabaseURL())
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	// get provider
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	googleOIDC, err := google.NewGoogleOIDC(ctx, config.GetGoogleClientID(), config.GetGoogleClientSecret(), google.CallbackURL)
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	// get service
	tsvc := services.NewTokenService(repo)
	asvc := services.NewAuthService(repo, googleOIDC, tsvc)
	usvc := services.NewUserService(repo)

	// make grpc server and listen
	grpcAdapter := grpc.NewAdapter(9001, asvc, usvc)

	if err := grpcAdapter.Run(); err != nil {
		log.Fatalf("error starting grpc server: %v\n", err)
	}
}
