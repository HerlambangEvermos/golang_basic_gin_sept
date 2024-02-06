package routes

import (
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetPosition(c *gin.Context) {
	positions := []models.Position{}

	// tanpa relational database
	// config.DB.Find(&positions)

	//dengan relational database
	config.DB.Preload("Department").Find(&positions)

	getPositionResponse := []models.GetPositionResponse{}

	for _, p := range positions {
		department := models.DepartmentResponse{
			ID:   p.Department.ID,
			Name: p.Department.Name,
			Code: p.Department.Code,
		}

		post := models.GetPositionResponse{
			ID:           p.ID,
			Name:         p.Name,
			Code:         p.Code,
			DepartmentID: p.DepartmentID,
			Department:   department,
		}

		getPositionResponse = append(getPositionResponse, post)
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Welcome Positions!",
		"data":    getPositionResponse,
	})
}

func PostPosition(c *gin.Context) {
	reqPos := models.Position{}
	c.BindJSON(&reqPos)

	config.DB.Create(&reqPos)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Insert Successfully",
		"data":    reqPos,
	})
}

func GetPositionById(c *gin.Context) {
	id := c.Param("id")

	position := models.Position{}

	// tanpa relationship
	// data := config.DB.First(&position, "id = ?", id)

	// dengan relationship
	data := config.DB.Preload("Department").First(&position, "id = ?", id)

	// validate data
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Position not found",
		})
		return
	}

	dept := models.DepartmentResponse{
		ID:   position.Department.ID,
		Name: position.Department.Name,
		Code: position.Department.Code,
	}

	getPositionResponse := models.GetPositionResponse{
		ID:           position.ID,
		Name:         position.Name,
		Code:         position.Code,
		DepartmentID: position.DepartmentID,
		Department:   dept,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    getPositionResponse,
	})
}

func PutPosition(c *gin.Context) {
	id := c.Param("id")

	reqPos := models.Position{}
	c.BindJSON(&reqPos)

	config.DB.Model(&models.Position{}).Where("id = ?", id).Updates(reqPos)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully",
		"data":    reqPos,
	})
}

func DeletePosition(c *gin.Context) {
	id := c.Param("id")

	position := models.Position{}

	config.DB.Delete(&position, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
