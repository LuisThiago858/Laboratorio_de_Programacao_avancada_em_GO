/*
Membros: Luis Thiago Silva Rabello
*/

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// serve a pagina estatica quando a porta 8080 e aberta
func caminhoPaginaEstatica(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// serve paginas dinamicas de acordo com a chave e valor passadas. Exemplo localhost:8080/info?key=goroutine
func paginaDinamica(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key") // pega a chave da URL
	response := ""

	//conteudo dinamico mostrado de acordo com o valor passado
	data := map[string]string{
		"golang":    "Golang é uma linguagem de programação criada pelo Google.",
		"goroutine": "Goroutine é uma função executada concorrentemente em Go.",
		"channel":   "Channel é um mecanismo de comunicação entre goroutines em Go.",
	}
	//caso o valor passado seja invalido o sistema mostra alguma opções validas
	if value, ok := data[strings.ToLower(key)]; ok {
		response = value
	} else {
		response = "Chave não encontrada. Tente 'golang', 'goroutine' ou 'channel"
	}

	//template basico das paginas dinamicas
	tmpl, _ := template.New("response").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Página Dinamica</title>
	</head>
	<body>
		<h1>Informação sobre Go</h1>
		<p>{{.}}</p>
	</body>
	</html>
	`)
	tmpl.Execute(w, response)

}

func main() {
	//passa a pagina estatica assim que o sistema e iniciado
	http.HandleFunc("/", caminhoPaginaEstatica)
	//de acordo com o valor da chave passado após info mostra o conteudo dinamico
	http.HandleFunc("/info", paginaDinamica)

	//mostra onde a porta foi aberta
	fmt.Println("Servidor aberto na porta 8080")
	//abre a porta
	http.ListenAndServe(":8080", nil)
}
