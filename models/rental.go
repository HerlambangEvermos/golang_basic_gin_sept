package models

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeInventory struct {
	gorm.Model
	EmployeeID  uint      `json:"employee_id"`
	Employee    Employee  `gorm:"foreignKey:EmployeeID;reference:ID"`
	InventoryID uint      `json:"inventory_id"`
	Inventory   Inventory `gorm:"foreignKey:InventoryID;reference:ID"`
	Description string    `json:"description"`
}

type ResponseGetRental struct {
	Id            uint      `json:"id"`
	Description   string    `json:"description"`
	EmployeeName  string    `json:"employeeName"`
	InventoryName string    `json:"inventoryName"`
	CreatedAt     time.Time `json:"createdAt"`
}

type ResponseEmployeeInventory struct {
	EmployeeID  uint      `json:"employee_id"`
	InventoryID uint      `json:"inventory_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type RequestRental struct {
	EmployeeID  uint   `json:"employee_id"`
	InventoryID uint   `json:"inventory_id"`
	Description string `json:"description"`
}

type ResponseRental struct {
	ID          uint      `json:"id"`
	EmployeeID  uint      `json:"employee_id"`
	InventoryID uint      `json:"inventory_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
