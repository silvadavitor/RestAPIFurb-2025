package main

import (
	"RestAPIFurb-2025/controller"
	"RestAPIFurb-2025/db"
	"RestAPIFurb-2025/repository"
	"RestAPIFurb-2025/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada repository
	ComandaRepository := repository.NewComandaRepository(dbConnection)

	//Camada usecase
	ComandaUseCase := usecase.NewComandaUsecase(ComandaRepository)

	//Camada controllers
	ComandaController := controller.NewComandaController(ComandaUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/RestAPIFurb/comandas", ComandaController.GetComandas)

	server.GET("/RestAPIFurb/comandas/:id", ComandaController.GetComandaById)

	server.POST("/RestAPIFurb/comandas", ComandaController.CreateComanda)

	server.PUT("/RestAPIFurb/comandas/:id", ComandaController.UpdateComanda)

	server.DELETE("/RestAPIFurb/comandas/:id", ComandaController.DeleteComanda)

	fmt.Println("Rodando servidor na porta :8080")

	server.Run(":8080")
}
