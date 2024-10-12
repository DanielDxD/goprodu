package main

import (
	"github.com/DanielDxD/goprodu/controller"
	"github.com/DanielDxD/goprodu/db"
	"github.com/DanielDxD/goprodu/repositories"
	"github.com/DanielDxD/goprodu/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	productRepository := repositories.NewProductRepository(dbConnection)

	productUseCase := usecases.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong",
		})
	})
	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)
	server.GET("/products/:productId", productController.GetProduct)

	server.Run(":8080")
}
