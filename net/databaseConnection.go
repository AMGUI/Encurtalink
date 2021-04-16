package net

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionData() (*sql.DB, error) {
	//user:senha@database
	strConexao := "root:admin@tcp(localhost:3306)/data_link"
	db, erro := sql.Open("mysql", strConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
