package usecase

import (
	"fmt"
	"tcg_pocket/model"
	"tcg_pocket/repository"
)

type TCGApoiadorUsecase struct {
	repository repository.TCGApoiadorRepository
}

func NewTCGApoiadorUseCase(repository repository.TCGApoiadorRepository) TCGApoiadorUsecase {
	return TCGApoiadorUsecase{repository}
}

func (u *TCGUseCase) GetTCGCollectionItem() ([]model.Item, error) {
	return u.repository.GetTCGCollectionItem()
}

func (u *TCGUseCase) GetTCGItemByID(id int) (model.Item, error) {
	return u.repository.GetTCGItemByID(id)
}

func (u *TCGUseCase) CreateItem(item model.Item) (model.Item, error) {
	id, err := u.repository.CreateItem(item)
	if err != nil {
		fmt.Printf("Erro ao tentar Criar o item: %v\n", err)
		return model.Item{}, err
	}
	item.Id = id
	return item, err
}

func (u *TCGUseCase) UpdateTCGItem(id int, model model.Item) (model.Item, error) {
	return u.repository.UpdateTCGItem(id, model)
}

func (u *TCGUseCase) DeleteTCGItem(id int) (string, error) {
	return u.repository.DeleteTCGItem(id)
}
