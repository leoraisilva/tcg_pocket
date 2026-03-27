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

	server.POST("/apoiador/create", controller.CreateApoiador)
	server.GET("/apoiador/:id", controller.GetTCGApoiadorByID)
	server.GET("/apoiadores", controller.GetTCGCollectionApoiador)

	server.Run(":8000")

}
