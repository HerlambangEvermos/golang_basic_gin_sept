package models

import "gorm.io/gorm"

type Archive struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	InventoryID uint   `json:"inventory_id"`
}

type GetArchiveResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
