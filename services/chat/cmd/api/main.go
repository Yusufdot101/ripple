package main

import (
	"log"

	"github.com/Yusufdot101/ripple/services/chat/config"
	"github.com/Yusufdot101/ripple/services/chat/internal/adapters/primary/api"
	"github.com/Yusufdot101/ripple/services/chat/internal/adapters/secondary/postgresql"
	"github.com/Yusufdot101/ripple/services/chat/internal/application/core/services"
)

func main() {
	repo, err := postgresql.NewAdapter(config.GetDatabaseURL())
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	csvc := services.NewChatService(repo)
	srv := api.NewServer(csvc)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}
}
