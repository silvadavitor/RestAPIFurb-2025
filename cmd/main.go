package main

import (
	"RestAPIFurb-2025/controller"
	"RestAPIFurb-2025/db"
	"RestAPIFurb-2025/middleware"
	"RestAPIFurb-2025/repository"
	"RestAPIFurb-2025/usecase"
	"fmt"
	"github.com/gin-gonic/gin"

	_ "RestAPIFurb-2025/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title REST API Comandas FURB
// @version 1.0
// @description Documentação Swagger da API de comandas da prova de suficiência Web II.
// @contact.name Vitor da Silva
// @contact.email vitsilva@furb.br
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada repository
	ComandaRepository := repository.NewComandaRepository(dbConnection)
	LoginRepository := repository.NewLoginRepository(dbConnection)

	//Camada usecase
	ComandaUseCase := usecase.NewComandaUsecase(ComandaRepository)
	LoginUseCase := usecase.NewLoginUsecase(LoginRepository)

	//Camada controllers
	ComandaController := controller.NewComandaController(ComandaUseCase)
	LoginController := controller.NewLoginController(LoginUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/RestAPIFurb/comandas", ComandaController.GetComandas)

	server.GET("/RestAPIFurb/comandas/:id", ComandaController.GetComandaById)

	server.POST("/RestAPIFurb/comandas", ComandaController.CreateComanda)

	server.PUT("/RestAPIFurb/comandas/:id", ComandaController.UpdateComanda)

	server.DELETE("/RestAPIFurb/comandas/:id", middleware.JWTAuthMiddleware(), ComandaController.DeleteComanda)

	server.POST("/RestAPIFurb/login", LoginController.Login)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("Rodando servidor na porta :8080")

	server.Run(":8080")
}
