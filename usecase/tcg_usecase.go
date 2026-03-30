package usecase

import (
	"fmt"
	"tcg_pocket/model"
	"tcg_pocket/repository"
)

type TCGUseCase struct {
	repository repository.TCGRepository
}

func NewTCGUseCase(repository repository.TCGRepository) TCGUseCase {
	return TCGUseCase{repository}
}

/* Endpoint /pokemon */
func (u *TCGUseCase) CreateTCGPokemon(model model.Pokemon) (model.Pokemon, error) {
	id, err := u.repository.CreateTCGPokemon(model)
	if err != nil {
		return model, err
	}
	model.Id = id
	return model, nil
}

func (u *TCGUseCase) GetTCGPokemonByID(id int) (model.Pokemon, error) {
	return u.repository.GetTCGPokemonByID(id)
}

func (u *TCGUseCase) GetTCGCollection() ([]model.Pokemon, error) {
	return u.repository.GetTCGCollection()
}

func (u *TCGUseCase) UpdateTCGPokemon(id int, model model.Pokemon) (model.Pokemon, error) {
	return u.repository.UpdateTCGPokemon(id, model)
}

func (u *TCGUseCase) DeleteTCGPokemon(id int) (string, error) {
	return u.repository.DeleteTCGPokemon(id)
}

/* Endpoint /apoiador */
func (u *TCGUseCase) CreateApoiador(apoiador model.Apoiador) (model.Apoiador, error) {
	id, err := u.repository.CreateApoiador(apoiador)
	if err != nil {
		fmt.Printf("Erro na criação do Apoiador: %v\n", err)
		return model.Apoiador{}, err
	}
	apoiador.Id = id
	return apoiador, err
}

func (u *TCGUseCase) GetTCGApoiadorByID(id int) (model.Apoiador, error) {
	return u.repository.GetTCGApoiadorByID(id)
}

func (u *TCGUseCase) GetTCGCollectionApoiador() ([]model.Apoiador, error) {
	return u.repository.GetTCGCollectionApoiador()
}

func (u *TCGUseCase) UpdateTCGApoiador(id int, model model.Apoiador) (model.Apoiador, error) {
	return u.repository.UpdateTCGApoiador(id, model)
}

func (u *TCGUseCase) DeleteTCGApoiador(id int) (string, error) {
	return u.repository.DeleteTCGApoiador(id)
}

/* Endpoint /item */
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
