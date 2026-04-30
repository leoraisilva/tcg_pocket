package repository

import (
	"database/sql"
	"fmt"
	"tcg_pocket/model"

	"github.com/lib/pq"
)

type TCGRepository struct {
	db *sql.DB
}

func NewTCGRepository(db *sql.DB) TCGRepository {
	return TCGRepository{db: db}
}

/* Endpoint /pokemon */
func (r *TCGRepository) CreateTCGPokemon(pokemon model.Pokemon) (int32, error) {
	var id int32

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Erro ao iniciar o Banco: %v\n", err)
		return 0, err
	}

	var ataqueAtualizado []model.Ataque

	for _, atk := range pokemon.Ataque {
		ataque, err := r.GetAtaqueForNome(tx, atk.Nome)
		if err != nil {
			if err == sql.ErrNoRows {
				novoAtaque, err := r.CreateAtaque(tx, atk)
				if err != nil {
					fmt.Printf("Erro ao criar ataque: %v\n", err)
					tx.Rollback()
					return 0, err
				}
				ataqueAtualizado = append(ataqueAtualizado, novoAtaque)
			} else {
				fmt.Printf("Erro ao buscar ataque: %v\n", err)
				tx.Rollback()
				return 0, err
			}
		} else {
			ataqueAtualizado = append(ataqueAtualizado, ataque)
		}
	}

	var habilidadeAtualizado []model.Habilidade

	for _, hab := range pokemon.Habilidade {
		habilidade, err := r.GetHabilidadeForNome(tx, hab.Nome)
		if err != nil {
			if err == sql.ErrNoRows {
				novaHabilidade, err := r.CreateHabilidade(tx, hab)
				if err != nil {
					fmt.Printf("Erro ao criar habilidade: %v\n", err)
					tx.Rollback()
					return 0, err
				}
				habilidadeAtualizado = append(habilidadeAtualizado, novaHabilidade)
			} else {
				fmt.Printf("Erro ao buscar habilidade: %v\n", err)
				tx.Rollback()
				return 0, err
			}
		} else {
			habilidadeAtualizado = append(habilidadeAtualizado, habilidade)
		}
	}

	err = tx.QueryRow(`
		INSERT INTO pokemon (nome, card_type, tipo, estagio, geracao, ps, recuo, fraqueza)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`,
		pokemon.Nome,
		pokemon.TipoCarta,
		pokemon.Tipo,
		pokemon.Estagio,
		pokemon.Geracao,
		pokemon.PS,
		pokemon.Recuo,
		pokemon.Fraqueza,
	).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar Pokemon: %v\n", err)
		tx.Rollback()
		return 0, err
	}

	for _, atk := range ataqueAtualizado {
		_, err := tx.Exec(`
			INSERT INTO pokemon_ataque (id_pokemon, ataque)
			VALUES ($1, $2)
		`, id, atk.Nome)

		if err != nil {
			fmt.Printf("Erro ao adicionar ataque: %v\n", err)
			tx.Rollback()
			return 0, err
		}
	}

	for _, hab := range habilidadeAtualizado {
		_, err := tx.Exec(`
			INSERT INTO pokemon_habilidade (id_pokemon, habilidade)
			VALUES ($1, $2)
		`, id, hab.Nome)

		if err != nil {
			fmt.Printf("Erro ao adicionar habilidade: %v\n", err)
			tx.Rollback()
			return 0, err
		}
	}

	return id, tx.Commit()
}

func (r *TCGRepository) CreateAtaque(tx *sql.Tx, ataque model.Ataque) (model.Ataque, error) {
	query := `INSERT INTO ataque (nome_ataque, dano_ataque, custo_ataque, efeito_ataque)
			  VALUES ($1, $2, $3, $4)`

	_, err := tx.Exec(query, ataque.Nome, ataque.Dano, pq.Array(ataque.Custo), ataque.Efeito)
	if err != nil {
		fmt.Printf("Erro ao criar ataque: %v\n", err)
		return ataque, err
	}
	return ataque, nil
}

func (r *TCGRepository) GetAtaqueForNome(tx *sql.Tx, nome string) (model.Ataque, error) {
	query := `SELECT nome_ataque, dano_ataque, custo_ataque, efeito_ataque FROM ataque WHERE nome_ataque = $1`
	var custoStr []string
	var ataque model.Ataque
	err := tx.QueryRow(query, nome).Scan(&ataque.Nome, &ataque.Dano, pq.Array(&custoStr), &ataque.Efeito)
	if err != nil {
		fmt.Printf("Erro ao buscar ataque: %v\n", err)
		return model.Ataque{}, err
	}
	for _, c := range custoStr {
		ataque.Custo = append(ataque.Custo, model.Tipo(c))
	}

	return ataque, err
}

func (r *TCGRepository) CreateHabilidade(tx *sql.Tx, ataque model.Habilidade) (model.Habilidade, error) {
	query := `INSERT INTO habilidade (nome_habilidade, efeito_habilidade) VALUES ($1, $2)`
	_, err := tx.Exec(query, ataque.Nome, ataque.Efeito)
	if err != nil {
		fmt.Printf("Erro ao criar ataque: %v\n", err)
		return ataque, err
	}
	return ataque, err
}

func (r *TCGRepository) GetHabilidadeForNome(tx *sql.Tx, nome string) (model.Habilidade, error) {
	query := `SELECT nome_habilidade, efeito_habilidade FROM habilidade WHERE nome_habilidade = $1`
	var habilidade model.Habilidade
	err := tx.QueryRow(query, nome).Scan(&habilidade.Nome, &habilidade.Efeito)
	if err != nil {
		fmt.Printf("Erro ao buscar habilidade: %v\n", err)
		return model.Habilidade{}, err
	}
	return habilidade, err
}

func (r *TCGRepository) GetTCGPokemonByID(id int32) (model.Pokemon, error) {
	var pokemon model.Pokemon
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Erro ao iniciar o Banco: %v\n", err)
		return model.Pokemon{}, err
	}
	query := `
	select
		id,
		nome,
		card_type,
		tipo,
		estagio,
		geracao,
		ps,
		recuo,
		fraqueza
	from pokemon where id=$1`
	err = tx.QueryRow(query, id).Scan(&pokemon.Id, &pokemon.Nome, &pokemon.TipoCarta, &pokemon.Tipo, &pokemon.Estagio, &pokemon.Geracao, &pokemon.PS, &pokemon.Recuo, &pokemon.Fraqueza)
	if err != nil {
		fmt.Printf("Erro ao buscar Pokemon: %v\n", err)
		return model.Pokemon{}, err
	}

	query = `select ataque from pokemon_ataque where id_pokemon=$1`
	ataque, err := tx.Query(query, id)
	if err != nil {
		fmt.Printf("Erro ao buscar ataque Pokemon: %v\n", err)
		return model.Pokemon{}, err
	}
	var listAtaque []string
	for ataque.Next() {
		var atk string
		if err = ataque.Scan(&atk); err != nil {
			fmt.Printf("Erro ao buscar ataque Pokemon: %v\n", err)
			return model.Pokemon{}, err
		}
		listAtaque = append(listAtaque, atk)
	}

	for _, atk := range listAtaque {
		valueAtaque, err := r.GetAtaqueForNome(tx, atk)
		if err != nil {
			fmt.Printf("Erro ao buscar ataque Pokemon: %v\n", err)
			return model.Pokemon{}, err
		}
		pokemon.Ataque = append(pokemon.Ataque, valueAtaque)
	}

	query = `select habilidade from pokemon_habilidade where id_pokemon=$1`
	habilidade, err := tx.Query(query, id)
	if err != nil {
		fmt.Printf("Erro ao buscar habilidade Pokemon: %v\n", err)
		return model.Pokemon{}, err
	}

	var listHabilidade []string
	for habilidade.Next() {
		var hab string
		if err = habilidade.Scan(&hab); err != nil {
			fmt.Printf("Erro ao buscar habilidade Pokemon: %v\n", err)
			return model.Pokemon{}, err
		}
		listHabilidade = append(listHabilidade, hab)
	}

	for _, hab := range listHabilidade {
		valueHabilidade, err := r.GetHabilidadeForNome(tx, hab)
		if err != nil {
			fmt.Printf("Erro ao buscar habilidade Pokemon: %v\n", err)
			return model.Pokemon{}, err
		}
		pokemon.Habilidade = append(pokemon.Habilidade, valueHabilidade)
	}

	return pokemon, tx.Commit()
}

func (r *TCGRepository) GetTCGCollection() ([]model.Pokemon, error) {

	query := `SELECT id FROM pokemon`
	list, err := r.db.Query(query)
	if err != nil {
		fmt.Printf("Erro ao Listar Pokemon: %v\n", err)
		return []model.Pokemon{}, err
	}
	var listPokemon []model.Pokemon
	var listId []int32
	for list.Next() {
		var isPokemon int32
		if err := list.Scan(&isPokemon); err != nil {
			fmt.Printf("Erro na listagem de Pokemon: %v\n", err)
			return []model.Pokemon{}, err
		}
		listId = append(listId, isPokemon)
	}

	for _, isPokemon := range listId {
		pokemon, err := r.GetTCGPokemonByID(isPokemon)
		if err != nil {
			fmt.Printf("Erro na listagem de Pokemon: %v\n", err)
			return []model.Pokemon{}, err
		}
		listPokemon = append(listPokemon, pokemon)
	}

	return listPokemon, err
}

func (r *TCGRepository) UpdateTCGPokemon(id int32, base model.Pokemon) (model.Pokemon, error) {
	var pokemon model.Pokemon

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("Erro ao iniciar o Banco: %v\n", err)
		return model.Pokemon{}, err
	}

	query := `UPDATE pokemon SET nome=$1, card_type=$2, tipo=$3, estagio=$4, geracao=$5, ps=$6, recuo=$7, fraqueza=$8 WHERE id=$9`
	err = tx.QueryRow(query, base.Nome, base.TipoCarta, base.Tipo, base.Estagio, base.Geracao, base.PS, base.Recuo, base.Fraqueza, id).Scan(
		&pokemon.Nome,
		&pokemon.TipoCarta,
		&pokemon.Tipo,
		&pokemon.Estagio,
		&pokemon.Geracao,
		&pokemon.PS,
		&pokemon.Recuo,
		&pokemon.Fraqueza,
	)
	if err != nil {
		fmt.Printf("Erro ao tentar alterar o pokemon: %v\n", err)
		tx.Rollback()
		return model.Pokemon{}, err
	}

	return pokemon, tx.Commit()
}

func (r *TCGRepository) DeleteTCGPokemon(id int32) (string, error) {
	response := "Delete com Sucesso!! "
	query := `DELETE FROM pokemon_ataque WHERE id_pokemon=$1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Erro ao desvincular ataque: %v\n", err)
		return "", err
	}
	query = `DELETE FROM pokemon_habilidade WHERE id_pokemon=$1`
	_, err = r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Erro ao desvincular habilidade: %v\n", err)
		return "", err
	}
	query = `DELETE FROM pokemon WHERE id=$1`
	_, err = r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Erro ao deletar o pokemon: %v\n", err)
		return "", err
	}

	return response, err
}
