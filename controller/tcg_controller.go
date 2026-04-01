package controller

import (
	"fmt"
	"tcg_pocket/model"
	"tcg_pocket/usecase"

	"github.com/gin-gonic/gin"
)

type TCGController struct {
	usecase usecase.TCGUseCase
}

func NewTCGController(usecase usecase.TCGUseCase) TCGController {
	return TCGController{usecase}
}

func (c *TCGController) CreateTCGPokemon(g *gin.Context) {
	var model model.Pokemon
	if err := g.ShouldBindJSON(&model); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := c.usecase.CreateTCGPokemon(model)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, result)
}

func (c *TCGController) GetTCGPokemonByID(g *gin.Context) {
	idParam := g.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		g.JSON(400, gin.H{"error": "ID inválido"})
		return
	}
	result, err := c.usecase.GetTCGPokemonByID(id)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, result)
}

func (c *TCGController) GetTCGCollection(g *gin.Context) {
	list, err := c.usecase.GetTCGCollection()
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, list)
}

func (c *TCGController) UpdateTCGPokemon(g *gin.Context) {
	var id int
	idParam := g.Param("id")
	_, err := fmt.Sscanf(idParam, "id", id)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var row model.Pokemon
	if err = g.ShouldBindJSON(&row); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := c.usecase.UpdateTCGPokemon(id, row)
	if err != nil {
		g.JSON(500, gin.H{"error": err})
		return
	}
	g.JSON(200, response)
}

func (c *TCGController) DeleteTCGPokemon(g *gin.Context) {
	var id int
	idParam := g.Param("id")
	id, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := c.usecase.DeleteTCGPokemon(id)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, response)
}
