package services

import InventoryClient "github.com/JacobRWebb/InventoryManagement.Inventory.Api/pkg"

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
