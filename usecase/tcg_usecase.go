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

func (u *TCGUseCase) CreateTCG(model model.Card) (model.Card, error) {
	id, err := u.repository.CreateTCG(model)
	if err != nil {
		return model, err
	}
	model.Id = id
	return model, nil
}
