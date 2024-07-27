package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/JacobRWebb/InventoryManagement.Inventory.Api/internal/config"
	"github.com/JacobRWebb/InventoryManagement.Inventory.Api/internal/database"
	"github.com/JacobRWebb/InventoryManagement.Inventory.Api/internal/server"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("there was an issue loading the config: %v", err)
	}

	db := database.MustOpen(cfg.DB.DSN)

	srv := server.NewServer(cfg, db)

	if err := srv.Run(); err != nil {
		log.Fatalf("there was an issue running the gRPC server: %v", err)
	}
}
