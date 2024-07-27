package services

import InventoryClient "github.com/JacobRWebb/InventoryManagement.Inventory.Api/pkg"

type inventoryService struct {
	InventoryClient.UnimplementedInventoryClientServer
}

func NewInventoryService() InventoryService {
	s := &inventoryService{}

	return s
}
