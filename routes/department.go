package routes

import (
	"fmt"
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

func GetDepartment(c *gin.Context) {
	departments := []models.Department{}
	// config.DB.Find(&departments)

	config.DB.Preload("Positions").Find(&departments)

	getDepartmentResponse := []models.GetDepartmentResponse{}

	for _, d := range departments {

		positions := []models.PositionResponse{}
		for _, p := range d.Positions {
			pos := models.PositionResponse{
				ID:   p.ID,
				Name: p.Name,
				Code: p.Code,
			}

			positions = append(positions, pos)
		}

		dept := models.GetDepartmentResponse{
			ID:        d.ID,
			Name:      d.Name,
			Code:      d.Code,
			Positions: positions,
		}

		getDepartmentResponse = append(getDepartmentResponse, dept)
	}

	t := time.Now()
	fmt.Println("time: ", t)

	fmt.Println("time: ", t.Format("02 Januari 2006"))

	c.JSON(http.StatusOK, gin.H{
		"Message": "Welcome Department!",
		"data":    getDepartmentResponse,
	})
}

func PostDepartment(c *gin.Context) {
	validate := validator.New()
	reqDep := models.Department{}
	c.BindJSON(&reqDep)

	errs := validate.Struct(&reqDep)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})

		c.Abort()
		return
	}

	config.DB.Create(&reqDep)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Insert Successfully",
		"data":    reqDep,
	})
}

func GetDepartmentById(c *gin.Context) {
	id := c.Param("id")

	department := models.Department{}
	// data := config.DB.First(&department, "id = ?", id)

	data := config.DB.Preload("Positions").First(&department, "id = ?", id)

	// validate data
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Department not found",
		})
		return
	}

	positions := []models.PositionResponse{}
	for _, p := range department.Positions {
		pos := models.PositionResponse{
			ID:   p.ID,
			Name: p.Name,
			Code: p.Code,
		}

		positions = append(positions, pos)
	}

	getDepartmentResponse := models.GetDepartmentResponse{
		ID:        department.ID,
		Name:      department.Name,
		Code:      department.Code,
		Positions: positions,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    getDepartmentResponse,
	})
}

func PutDepartment(c *gin.Context) {
	id := c.Param("id")

	reqDep := models.Department{}
	c.BindJSON(&reqDep)

	config.DB.Model(&models.Department{}).Where("id = ?", id).Updates(reqDep)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully",
		"data":    reqDep,
	})
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	department := models.Department{}

	config.DB.Delete(&department, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
