package helper

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "tcg_db"
)

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = ConnMigration(db)
	if err != nil {
		return nil, err
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return db, nil
}

func ConnMigration(db *sql.DB) error {
	sqlByte, err := os.ReadFile("resource/migration.sql")
	if err != nil {
		fmt.Printf("Erro ao acessar as tabelas do banco: %v\n", err)
		return err
	}

	_, err = db.Exec(string(sqlByte))
	if err != nil {
		fmt.Printf("Erro ao acessar as tabelas do banco: %v\n", err)
		return err
	}
	fmt.Println("Migrations executado com sucesso!!")
	return nil
}
