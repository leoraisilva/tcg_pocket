package controller

import (
	"fmt"
	"tcg_pocket/model"
	"tcg_pocket/usecase"

	"github.com/gin-gonic/gin"
)

type TCGItemController struct {
	usecase usecase.TCGItemUseCase
}

func NewTCGItemController(usecase usecase.TCGItemUseCase) TCGItemController {
	return TCGItemController{usecase: usecase}
}

func (c *TCGController) GetTCGCollectionItem(g *gin.Context) {
	list, err := c.usecase.GetTCGCollectionItem()
	if err != nil {
		g.JSON(500, gin.H{"erro": err.Error()})
		return
	}
	g.JSON(200, list)
}

func (c *TCGController) GetTCGItemByID(g *gin.Context) {
	idParam := g.Param("id")
	var id int32
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := c.usecase.GetTCGItemByID(id)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, response)
}

func (c *TCGController) CreateItem(g *gin.Context) {
	var item model.Item
	if err := g.ShouldBindJSON(&item); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := c.usecase.CreateItem(item)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, response)
}

func (c *TCGController) UpdateTCGItem(g *gin.Context) {
	var id int32
	idParam := g.Param("id")
	_, err := fmt.Sscanf(idParam, "id", id)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var row model.Item
	if err = g.ShouldBindJSON(&row); err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := c.usecase.UpdateTCGItem(id, row)
	if err != nil {
		g.JSON(500, gin.H{"error": err})
		return
	}
	g.JSON(200, response)
}

func (c *TCGController) DeleteTCGItem(g *gin.Context) {
	idParam := g.Param("id")
	var id int32
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := c.usecase.DeleteTCGItem(id)
	if err != nil {
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}
	g.JSON(200, response)
}
