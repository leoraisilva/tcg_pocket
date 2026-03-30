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

func (r *TCGRepository) CreateTCGPokemon(model model.Pokemon) (int, error) {
	var id int

	ataque, err := r.GetAtaqueForNome(model.Ataque.Nome)
	if err != nil {
		if err == sql.ErrNoRows {
			novoAtaque, err := r.CreateAtaque(model.Ataque)
			if err != nil {
				fmt.Printf("Erro ao criar ataque: %v\n", err)
				return 0, err
			}
			model.Ataque = novoAtaque
		} else {
			fmt.Printf("Erro ao buscar ataque: %v\n", err)
			return 0, err
		}
	} else {
		model.Ataque = ataque
	}

	habilidade, err := r.GetHabilidadeForNome(model.Habilidade.Nome)
	if err != nil {
		if err == sql.ErrNoRows {
			novaHabilidade, err := r.CreateHabilidade(model.Habilidade)
			if err != nil {
				fmt.Printf("Erro ao criar habilidade: %v\n", err)
				return 0, err
			}
			model.Habilidade = novaHabilidade
		} else {
			fmt.Printf("Erro ao buscar habilidade: %v\n", err)
			return 0, err
		}
	} else {
		model.Habilidade = habilidade
	}

	err = r.db.QueryRow(`
		INSERT INTO pokemon (nome, card_type, tipo, estagio, habilidade, ataque, ps, recuo, fraqueza)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`,
		model.Nome,
		model.TipoCarta,
		model.Tipo,
		model.Estagio,
		model.Habilidade.Nome,
		model.Ataque.Nome,
		model.PS,
		model.Recuo,
		model.Fraqueza,
	).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar Pokemon: %v\n", err)
		return 0, err
	}

	return id, nil
}

func (r *TCGRepository) CreateAtaque(ataque model.Ataque) (model.Ataque, error) {
	query := `INSERT INTO ataque (nome_ataque, dano_ataque, custo_ataque, efeito_ataque) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, ataque.Nome, ataque.Dano, ataque.Custo, ataque.Efeito)
	if err != nil {
		fmt.Printf("Erro ao criar ataque: %v\n", err)
		return ataque, err
	}
	return ataque, err
}

func (r *TCGRepository) GetAtaqueForNome(nome string) (ataque model.Ataque, err error) {
	query := `SELECT nome_ataque, dano_ataque, custo_ataque, efeito_ataque FROM ataque WHERE nome_ataque = $1`
	if err != nil {
		fmt.Printf("Erro ao preparar query: %v\n", err)
		return model.Ataque{}, err
	}
	err = r.db.QueryRow(query, nome).Scan(&ataque.Nome, &ataque.Dano, &ataque.Custo, &ataque.Efeito)
	if err != nil {
		fmt.Printf("Erro ao buscar ataque: %v\n", err)
		return model.Ataque{}, err
	}
	return ataque, err
}

func (r *TCGRepository) CreateHabilidade(ataque model.Habilidade) (model.Habilidade, error) {
	query := `INSERT INTO habilidade (nome_habilidade, efeito_habilidade) VALUES ($1, $2)`
	_, err := r.db.Exec(query, ataque.Nome, ataque.Efeito)
	if err != nil {
		fmt.Printf("Erro ao criar ataque: %v\n", err)
		return ataque, err
	}
	return ataque, err
}

func (r *TCGRepository) GetHabilidadeForNome(nome string) (habilidade model.Habilidade, err error) {
	query := `SELECT nome_habilidade, efeito_habilidade FROM habilidade WHERE nome_habilidade = $1`
	if err != nil {
		fmt.Printf("Erro ao preparar query: %v\n", err)
		return model.Habilidade{}, err
	}
	err = r.db.QueryRow(query, nome).Scan(&habilidade.Nome, &habilidade.Efeito)
	if err != nil {
		fmt.Printf("Erro ao buscar habilidade: %v\n", err)
		return model.Habilidade{}, err
	}
	return habilidade, err
}

func (r *TCGRepository) GetTCGPokemonByID(id int) (model.Pokemon, error) {
	var pokemon model.Pokemon
	query := `SELECT id, nome, card_type, tipo, estagio, habilidade, ataque, ps, recuo, fraqueza FROM pokemon WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&pokemon.Id, &pokemon.Nome, &pokemon.TipoCarta, &pokemon.Tipo, &pokemon.Estagio, &pokemon.Habilidade.Nome, &pokemon.Ataque.Nome, &pokemon.PS, &pokemon.Recuo, &pokemon.Fraqueza)
	if err != nil {
		fmt.Printf("Erro ao buscar Pokemon: %v\n", err)
		return model.Pokemon{}, err
	}

	return pokemon, nil
}

func (r *TCGRepository) GetTCGCollection() ([]model.Pokemon, error) {

	query := `SELECT id, nome, card_type, tipo, estagio, habilidade, ataque, ps, recuo, fraqueza FROM pokemon`
	list, err := r.db.Query(query)
	if err != nil {
		fmt.Printf("Erro ao Listar Pokemon: %v\n", err)
		return []model.Pokemon{}, err
	}
	var listPokemon []model.Pokemon
	var pokemon model.Pokemon

	for list.Next() {
		err = list.Scan(
			&pokemon.Id,
			&pokemon.Nome,
			&pokemon.TipoCarta,
			&pokemon.Tipo,
			&pokemon.Estagio,
			&pokemon.Habilidade.Nome,
			&pokemon.Ataque.Nome,
			&pokemon.PS,
			&pokemon.Recuo,
			&pokemon.Fraqueza,
		)

		if err != nil {
			fmt.Printf("Erro na Listagem Pokemon: %v\n", err)
			return []model.Pokemon{}, err
		}
		listPokemon = append(listPokemon, pokemon)
	}
	list.Close()
	return listPokemon, err
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
