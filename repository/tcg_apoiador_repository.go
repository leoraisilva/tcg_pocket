package repository

import (
	"database/sql"
	"fmt"
	"tcg_pocket/model"
)

type TCGItemRepository struct {
	db *sql.DB
}

func NewTCGItemRepository(db *sql.DB) TCGItemRepository {
	return TCGItemRepository{db: db}
}

func (r *TCGRepository) CreateApoiador(apoiador model.Apoiador) (int, error) {
	var id int
	query := `INSERT INTO apoiador (nome, card_type, efeito) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, apoiador.Nome, apoiador.CardType, apoiador.Efeito).Scan(&id)
	if err != nil {
		fmt.Printf("Erro ao Criar Apoiador: %v\n", err)
		return 0, err
	}
	return id, err
}

func (r *TCGRepository) GetTCGApoiadorByID(id int) (model.Apoiador, error) {
	var apoiador model.Apoiador
	query := `SELECT id, nome, card_type, efeito FROM apoiador WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&apoiador.Id,
		&apoiador.Nome,
		&apoiador.CardType,
		&apoiador.Efeito)
	if err != nil {
		fmt.Printf("Erro ao Buscar Apoiador: %v\n", err)
		return model.Apoiador{}, err
	}
	return apoiador, err
}

func (r *TCGRepository) GetTCGCollectionApoiador() ([]model.Apoiador, error) {
	query := `SELECT id, nome, card_type, efeito FROM apoiador`
	list, err := r.db.Query(query)
	if err != nil {
		fmt.Printf("Erro ao listar os Apoiadores: %v\n", err)
		return []model.Apoiador{}, err
	}
	var apoiador model.Apoiador
	var collectioApoiador []model.Apoiador

	for list.Next() {
		err = list.Scan(
			&apoiador.Id,
			&apoiador.Nome,
			&apoiador.CardType,
			&apoiador.Efeito)
		if err != nil {
			fmt.Printf("Erro ao Listar Apoiadores: %v\n", err)
			return []model.Apoiador{}, err
		}
		collectioApoiador = append(collectioApoiador, apoiador)
	}
	list.Close()
	return collectioApoiador, err
}

func (r *TCGRepository) UpdateTCGApoiador(id int, apoiador model.Apoiador) (model.Apoiador, error) {
	var row model.Apoiador
	query := `UPDATE apoiador SET nome=$1, card_type=$2, efeito=$3 WHERE id=$4`
	err := r.db.QueryRow(query, apoiador.Nome, apoiador.CardType, apoiador.Efeito, id).Scan(
		&row.Nome,
		&row.CardType,
		&row.Efeito,
	)
	if err != nil {
		fmt.Printf("Erro ao tentar alterar o apoiador: %v\n", err)
		return model.Apoiador{}, err
	}
	return model.Apoiador{}, err
}

func (r *TCGRepository) DeleteTCGApoiador(id int) (string, error) {
	response := "Delete com Sucesso!!"
	query := `DELETE FROM apoiador WHERE id=$1`
	_, err := r.db.Query(query, id)
	if err != nil {
		fmt.Printf("Erro ao tentar deletar o Apoidores: %v\n", err)
		return "", err
	}
	return response, err
}
