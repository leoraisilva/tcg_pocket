package usecase

import (
	"fmt"
	"tcg_pocket/model"
	"tcg_pocket/repository"
)

type TCGItemUseCase struct {
	repository repository.TCGItemRepository
}

func NewTCGItemUseCase(repository repository.TCGItemRepository) TCGItemUseCase {
	return TCGItemUseCase{repository}
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

func (u *TCGUseCase) GetTCGApoiadorByID(id int32) (model.Apoiador, error) {
	return u.repository.GetTCGApoiadorByID(id)
}

func (u *TCGUseCase) GetTCGCollectionApoiador() ([]model.Apoiador, error) {
	return u.repository.GetTCGCollectionApoiador()
}

func (u *TCGUseCase) UpdateTCGApoiador(id int32, model model.Apoiador) (model.Apoiador, error) {
	return u.repository.UpdateTCGApoiador(id, model)
}

func (u *TCGUseCase) DeleteTCGApoiador(id int32) (string, error) {
	return u.repository.DeleteTCGApoiador(id)
}
