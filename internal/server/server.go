package server

import (
	"fmt"
	"log"
	"net"

	"github.com/JacobRWebb/InventoryManagement.Inventory.Api/internal/config"
	"github.com/JacobRWebb/InventoryManagement.Inventory.Api/internal/services"
	InventoryClient "github.com/JacobRWebb/InventoryManagement.Inventory.Api/pkg"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	cfg        config.Config
	gRPCServer *grpc.Server
}

func NewServer(cfg config.Config, db *gorm.DB) *Server {
	gRPCServer := grpc.NewServer()

	return &Server{
		cfg:        cfg,
		gRPCServer: gRPCServer,
	}
}

func (s *Server) Run() error {
	return s.startGrpcServer()
}

func (s *Server) startGrpcServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Service.Port))

	if err != nil {
		return err
	}

	services := services.NewServices()

	InventoryClient.RegisterInventoryClientServer(s.gRPCServer, services.InventoryService)

	log.Println("running gRPC server")

	return s.gRPCServer.Serve(lis)
}
