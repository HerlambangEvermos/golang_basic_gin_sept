package main

import (
	"golang_basic_gin_sept_2023/config"
	"golang_basic_gin_sept_2023/midlewares"
	"golang_basic_gin_sept_2023/routes"

	"github.com/gin-gonic/gin"
)

// commit ke dua

func main() {
	config.InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/home", GetHome)

	route := r.Group("/")
	{

		user := route.Group("/user")
		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.GenerateToken)
		}

		department := route.Group("departments").Use(midlewares.IsAdmin())
		{
			department.GET("/", routes.GetDepartment)
			department.GET("/:id", routes.GetDepartmentById)
			department.POST("/", routes.PostDepartment)
			department.PUT("/:id", routes.PutDepartment)
			department.DELETE("/:id", routes.DeleteDepartment)
		}

		position := route.Group("positions").Use(midlewares.IsAdmin())
		{
			position.GET("/", routes.GetPosition)
			position.GET("/:id", routes.GetPositionById)
			position.POST("/", routes.PostPosition)
			position.PUT("/:id", routes.PutPosition)
			position.DELETE("/:id", routes.DeletePosition)
		}

		inventory := route.Group("inventory").Use(midlewares.IsAdmin())
		{
			inventory.GET("/", routes.GetInventory)
			inventory.GET("/:id", routes.GetInventoryById)
			inventory.POST("/", routes.PostInventory)
			inventory.PUT("/:id", routes.PutInventory)
			inventory.DELETE("/:id", routes.DeleteInventory)
		}
		employees := route.Group("employees").Use(midlewares.Auth())
		{
			employees.GET("/", routes.GetEmployees)
			employees.GET("/:id", routes.GetEmployeesByID)
			employees.POST("/", routes.PostEmployees)
			employees.PUT("/:id", routes.PutEmployees)
			employees.DELETE("/:id", routes.DeleteEmployees)
		}
		rental := route.Group("rental").Use(midlewares.Auth())
		{
			rental.GET("/", routes.GetRental)
			rental.POST("/employee", routes.RentalByEmployeeID)
			rental.GET("/inventory/:id", routes.GetRentalByInventoryID)
		}
	}

	r.Run(":9001") // listen and serve on 0.0.0.0:8080
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Welcome Home!",
	})
}
