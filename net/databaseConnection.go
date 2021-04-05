package net

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionData() {
	//user:senha@database
	strConexao := "root:bibi2019@tcp(172.17.0.2:3306)/data_link"
	db, erro := sql.Open("mysql", strConexao)
	if erro != nil {
		fmt.Println("erro da conexer!")
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		fmt.Println("erro senha!")
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")
	/*
		linhas, erro := db.Query("select * from Dados_URL")

		if erro != nil {
			log.Fatal(erro)
		}
		defer linhas.Close()

		fmt.Println(linhas)
	*/
}
