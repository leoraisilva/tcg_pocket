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
