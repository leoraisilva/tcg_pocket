package repository

import (
	"database/sql"
	"fmt"
	"tcg_pocket/model"
)

type TCGRepository struct {
	db *sql.DB
}

func NewTCGRepository(db *sql.DB) TCGRepository {
	return TCGRepository{db: db}
}

func (r *TCGRepository) CreateTCG(model model.Card) (int, error) {

	var id int
	query, err := r.db.Prepare(`INSERT INTO card (nome, tipo, estagio, habilidade, ataque, ps, recuo, fraqueza) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`)
	if err != nil {
		fmt.Printf("Erro ao preparar query: %v\n", err)
		return 0, err
	}
	err = query.QueryRow(query, model.Nome, model.Tipo, model.Estagio, model.Habilidade, model.Ataque, model.PS, model.Recuo, model.Fraqueza).Scan(&id)
	if err != nil {
		fmt.Printf("Erro ao criar card: %v\n", err)
		return 0, err
	}
	return id, err
}

func (r *TCGRepository) CreateAtaque(ataque model.Ataque) (model.Ataque, error) {
	query := `INSERT INTO ataque (nome, dano, efeito) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, ataque.Nome, ataque.Dano, ataque.Efeito)
	if err != nil {
		fmt.Printf("Erro ao criar ataque: %v\n", err)
		return ataque, err
	}
	return ataque, err
}

func (r *TCGRepository) GetAtaqueForNome(nome string) (ataque model.Ataque, err error) {
	query := `SELECT nome, dano, efeito FROM ataque WHERE nome = $1`
	row := r.db.QueryRow(query, nome)
	err = row.Scan(&ataque.Nome, &ataque.Dano, &ataque.Efeito)
	return ataque, err
}
