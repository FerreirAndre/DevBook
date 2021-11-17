package banco

import (
	"api/src/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //Driver
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		log.Fatal(erro)
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
