package repository

import (
	"database/sql"
	"fmt"
	"tcg_pocket/model"
)

type TCGApoiadorRepository struct {
	db *sql.DB
}

func NewTCGApoiadorRepository(db *sql.DB) TCGApoiadorRepository {
	return TCGApoiadorRepository{db: db}
}

func (r *TCGRepository) CreateItem(item model.Item) (int, error) {
	var id int
	query := `INSERT INTO item (nome, card_type, efeito) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, item.Nome, item.CardType, item.Efeito).Scan(&id)
	if err != nil {
		fmt.Printf("Erro ao Criar Item: %v\n", err)
		return 0, err
	}
	return id, err
}

func (r *TCGRepository) GetTCGItemByID(id int) (model.Item, error) {
	var item model.Item
	query := `SELECT id, nome, card_type, efeito FROM item WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(
		&item.Id,
		&item.Nome,
		&item.CardType,
		&item.Efeito,
	)
	if err != nil {
		fmt.Printf("Erro ao Buscar Item por ID: %v\n", err)
		return model.Item{}, err
	}
	return item, err
}

func (r *TCGRepository) GetTCGCollectionItem() ([]model.Item, error) {
	query := `SELECT id, nome, card_type, efeito FROM item`
	list, err := r.db.Query(query)
	var item model.Item
	var itens []model.Item

	for list.Next() {
		list.Scan(
			&item.Id,
			&item.Nome,
			&item.CardType,
			&item.Efeito,
		)
		if err != nil {
			fmt.Printf("Erro ao escanear os item listado: %v\n", err)
			return []model.Item{}, err
		}
		itens = append(itens, item)
	}
	list.Close()
	return itens, err
}

func (r *TCGRepository) UpdateTCGItem(id int, item model.Item) (model.Item, error) {
	var row model.Item
	query := `UPDATE item SET nome=$1, card_type=$2, efeito=$3 WHERE id=$4`
	err := r.db.QueryRow(query, item.Nome, item.CardType, item.Efeito, id).Scan(
		&row.Nome,
		&row.CardType,
		&row.Efeito,
	)
	if err != nil {
		fmt.Printf("Erro ao tentar alterar o apoiador: %v\n", err)
		return model.Item{}, err
	}
	return model.Item{}, err
}

func (r *TCGRepository) DeleteTCGItem(id int) (string, error) {
	response := "Delete com Sucesso!!"
	_, err := r.db.Query(`DELETE FROM item WHERE id=$1`, id)
	if err != nil {
		fmt.Printf("Erro ao tentar deletar item: %v\n", err)
		return "", err
	}
	return response, err
}
