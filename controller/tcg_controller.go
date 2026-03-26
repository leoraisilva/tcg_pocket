package controller

import (
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

func (c *TCGController) CreateTCG(g *gin.Context) {
	var model model.Card
	if err := g.ShouldBindJSON(&model); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := c.usecase.CreateTCG(model)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, result)
}
