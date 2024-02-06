package routes

import (
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetInventory(c *gin.Context) {
	inventory := []models.Inventory{}

	//dengan relational database
	config.DB.Preload("Archive").Find(&inventory)

	getInventoryResponse := []models.GetInventoryResponse{}

	for _, i := range inventory {
		inv := models.GetInventoryResponse{
			InventoryID:          i.ID,
			InventoryName:        i.Name,
			InventoryDescription: i.Description,
			ArchiveName:          i.Archive.Name,
			ArchiveDescription:   i.Archive.Description,
		}

		getInventoryResponse = append(getInventoryResponse, inv)
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Welcome Inventorys!",
		"data":    getInventoryResponse,
	})
}

func PostInventory(c *gin.Context) {
	reqInv := models.RequestInventory{}
	c.BindJSON(&reqInv)

	inventory := models.Inventory{
		Name:        reqInv.InventoryName,
		Description: reqInv.InventoryDescription,
		Archive: models.Archive{
			Name:        reqInv.ArchiveName,
			Description: reqInv.ArchiveDescription,
		},
	}

	config.DB.Create(&inventory)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Insert Successfully",
		"data":    inventory,
	})
}

func GetInventoryById(c *gin.Context) {
	id := c.Param("id")

	Inventory := models.Inventory{}

	// dengan relationship
	data := config.DB.Preload("Archive").First(&Inventory, "id = ?", id)

	// validate data
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Inventory not found",
		})
		return
	}

	inv := models.GetInventoryResponse{
		InventoryID:          Inventory.ID,
		InventoryName:        Inventory.Name,
		InventoryDescription: Inventory.Description,
		ArchiveName:          Inventory.Archive.Name,
		ArchiveDescription:   Inventory.Archive.Description,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    inv,
	})
}

func PutInventory(c *gin.Context) {
	id := c.Param("id")

	data := config.DB.First(&models.Inventory{}, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	reqInv := models.RequestInventory{}
	c.BindJSON(&reqInv)

	inv := models.Inventory{
		Name:        reqInv.InventoryName,
		Description: reqInv.InventoryDescription,
	}
	config.DB.Model(&models.Inventory{}).Where("id = ?", id).Updates(&inv)

	archive := models.Archive{
		Name:        reqInv.ArchiveName,
		Description: reqInv.ArchiveDescription,
	}
	config.DB.Model(&models.Archive{}).Where("inventory_id = ?", id).Updates(&archive)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully",
		"data":    inv,
	})
}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")

	Inventory := models.Inventory{}

	data := config.DB.First(&Inventory, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	config.DB.Delete(&Inventory, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
