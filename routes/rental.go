package routes

import (
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetRental(c *gin.Context) {
	EmployeeInventory := []models.EmployeeInventory{}

	config.DB.Preload(clause.Associations).Find(&EmployeeInventory)

	responseGetRentals := []models.ResponseGetRental{}

	for _, ei := range EmployeeInventory {
		rgr := models.ResponseGetRental{
			Id:            ei.ID,
			Description:   ei.Description,
			EmployeeName:  ei.Employee.Name,
			InventoryName: ei.Inventory.Name,
			CreatedAt:     ei.CreatedAt,
		}

		responseGetRentals = append(responseGetRentals, rgr)
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "GET Rental Successfully",
		"data":    responseGetRentals,
	})

}

func RentalByEmployeeID(c *gin.Context) {
	var reqRental models.RequestRental

	if err := c.ShouldBindJSON(&reqRental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	rental := models.EmployeeInventory{
		EmployeeID:  reqRental.EmployeeID,
		InventoryID: reqRental.InventoryID,
		Description: reqRental.Description,
	}

	insert := config.DB.Create(&rental)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   insert.Error.Error(),
		})

		c.Abort()
		return
	}

	respRental := models.ResponseRental{
		ID:          rental.ID,
		EmployeeID:  rental.EmployeeID,
		InventoryID: rental.InventoryID,
		Description: rental.Description,
		CreatedAt:   rental.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Insert Rental Successfully",
		"data":    respRental,
	})

}

func GetRentalByInventoryID(c *gin.Context) {
	id := c.Param("id")

	inventories := models.Inventory{}
	emInv := []models.ResponseEmployeeInventory{}

	config.DB.Preload(clause.Associations).First(&inventories, "id = ?", id)

	for _, inv := range inventories.Employees {

		emInv = append(emInv, models.ResponseEmployeeInventory{
			EmployeeID:  inv.EmployeeID,
			InventoryID: inv.InventoryID,
			Description: inv.Description,
			CreatedAt:   inv.CreatedAt,
		})

	}

	respInv := models.ResponseInventoryEmployee{
		InventoryName:        inventories.Name,
		InventoryDescription: inventories.Description,
		EmployeeInventory:    emInv,
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Rental Data By Inventory ID!",
		"data":    respInv,
	})

}
