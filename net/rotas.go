package net

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"modulo/link"
	"net/http"
)

var templates *template.Template

//Inicail rota que chama o html
func Inicial(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("*.html"))

	templates.ExecuteTemplate(w, "home.html", nil)

}

//PostiLink funcao para utiliza do metodo post para inserir json no banco
func PostLink(w http.ResponseWriter, r *http.Request) {

	//obtendo Corpo da requisição

	bodyrequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var bodyJson link.ModeJason
	if erro = json.Unmarshal(bodyrequisicao, &bodyJson); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	PostBanco(link.Datamode(bodyJson.URL))

}
