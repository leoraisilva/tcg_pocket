package usecase

import (
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

func (u *TCGUseCase) UpdateTCGPokemon(id int, model model.Pokemon) (model.Pokemon, error) {
	return u.repository.UpdateTCGPokemon(id, model)
}

func (u *TCGUseCase) DeleteTCGPokemon(id int) (string, error) {
	return u.repository.DeleteTCGPokemon(id)
}
