package controller

import (
	"fmt"
	"tcg_pocket/model"
	"tcg_pocket/usecase"

	"github.com/gin-gonic/gin"
)

type TCGApoiadorController struct {
	usecase usecase.TCGApoiadorUsecase
}

func NewTCGApoiadorController(usecase usecase.TCGApoiadorUsecase) TCGApoiadorController {
	return TCGApoiadorController{usecase: usecase}
}

func (c *TCGController) CreateApoiador(g *gin.Context) {
	var apoiador model.Apoiador
	if err := g.ShouldBindJSON(&apoiador); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
	}
	response, err := c.usecase.CreateApoiador(apoiador)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
	}
	g.JSON(200, response)

}

func (c *TCGController) GetTCGApoiadorByID(g *gin.Context) {
	idParam := g.Param("id")
	var id int32
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := c.usecase.GetTCGApoiadorByID(id)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, response)
}

func (c *TCGController) GetTCGCollectionApoiador(g *gin.Context) {
	list, err := c.usecase.GetTCGCollectionApoiador()
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, list)
}

func (c *TCGController) UpdateTCGApoiador(g *gin.Context) {
	var id int32
	idParam := g.Param("id")
	_, err := fmt.Sscanf(idParam, "id", id)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var row model.Apoiador
	if err = g.ShouldBindJSON(&row); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := c.usecase.UpdateTCGApoiador(id, row)
	if err != nil {
		g.JSON(500, gin.H{"error": err})
		return
	}
	g.JSON(200, response)
}

func (c *TCGController) DeleteTCGApoiador(g *gin.Context) {
	idParam := g.Param("id")
	var id int32
	_, err := fmt.Sscanf(idParam, "id", &id)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := c.usecase.DeleteTCGApoiador(id)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, response)
}
