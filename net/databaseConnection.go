package net

import (
	"database/sql"
	"fmt"
	"modulo/link"

	_ "github.com/go-sql-driver/mysql"
)

func connectionData() (*sql.DB, error) {
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

func PostBanco(dadoParaBanco link.ModeJason) {

	db, err := connectionData()
	if err != nil {
		fmt.Println("Erro ao conectar no banco de dados!")
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into dados_url (linkUrl, linkmenor) values (?, ?)")
	if err != nil {
		fmt.Println("Erro no statement !")
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(dadoParaBanco.URL, dadoParaBanco.LinkMenor)
	if err != nil {
		fmt.Println("Erro ao inserir no banco de dados!")
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao retorna ID!")
		return
	}

	fmt.Printf("Link inserido com sucesso! Id: %d", idInserido)

}
