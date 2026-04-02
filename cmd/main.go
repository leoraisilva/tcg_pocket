package main

import (
	"tcg_pocket/controller"
	"tcg_pocket/helper"
	"tcg_pocket/repository"
	"tcg_pocket/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	conn, err := helper.GetConnection()
	if err != nil {
		panic(err)
	}

	repository := repository.NewTCGRepository(conn)
	usecase := usecase.NewTCGUseCase(repository)
	controller := controller.NewTCGController(usecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"mensagem": "Servidor rodando com sucesso",
		})
	})

	server.POST("/pokemon/create", controller.CreateTCGPokemon)
	server.GET("/pokemon/:id", controller.GetTCGPokemonByID)
	server.GET("/pokemons", controller.GetTCGCollection)
	server.PUT("/pokemon/update/:id", controller.UpdateTCGPokemon)
	server.DELETE("/pokemon/delete/:id", controller.DeleteTCGPokemon)

	server.POST("/apoiador/create", controller.CreateApoiador)
	server.GET("/apoiador/:id", controller.GetTCGApoiadorByID)
	server.GET("/apoiadores", controller.GetTCGCollectionApoiador)
	server.PUT("/apoiador/update/:id", controller.UpdateTCGApoiador)
	server.DELETE("/apoiador/delete/:id", controller.DeleteTCGApoiador)

	server.POST("/item/create", controller.CreateItem)
	server.GET("/item/:id", controller.GetTCGItemByID)
	server.GET("/itens", controller.GetTCGCollectionItem)
	server.PUT("/item/update/:id", controller.UpdateTCGItem)
	server.DELETE("/item/delete/:id", controller.DeleteTCGItem)

	server.Run(":8000")

}
