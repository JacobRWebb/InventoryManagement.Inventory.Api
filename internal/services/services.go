package services

import InventoryClient "github.com/JacobRWebb/InventoryManagement.Inventory.Api/internal/protos/inventory"

type Services struct {
	InventoryService InventoryService
}

type InventoryService interface {
	InventoryClient.InventoryClientServer
}

func NewServices() *Services {
	clients := &Services{
		InventoryService: NewInventoryService(),
	}

	return clients
}
