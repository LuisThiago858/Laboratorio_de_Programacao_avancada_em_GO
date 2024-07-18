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

func caminhoPaginaEstatica(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func paginaDinamica(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")
	response := ""

	data := map[string]string{
		"golang":    "Golang é uma linguagem de programação criada pelo Google.",
		"goroutine": "Goroutine é uma função executada concorrentemente em Go.",
		"channel":   "Channel é um mecanismo de comunicação entre goroutines em Go.",
	}

	if value, ok := data[strings.ToLower(key)]; ok {
		response = value
	} else {
		response = "Chave não encontrada. Tente 'golang', 'goroutine' ou 'channel"
	}

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
	http.HandleFunc("/", caminhoPaginaEstatica)
	http.HandleFunc("/info", paginaDinamica)

	fmt.Println("Servidor aberto na porta 8080")
	http.ListenAndServe(":8080", nil)
}
