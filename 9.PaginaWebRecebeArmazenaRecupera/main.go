package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// Função principal onde o servidor web é iniciado
func main() {
	http.HandleFunc("/", indexHandler)            // Manipulador para a página principal
	http.HandleFunc("/save", saveHandler)         // Manipulador para salvar os dados
	http.HandleFunc("/retrieve", retrieveHandler) // Manipulador para recuperar os dados
	fmt.Println("Server esta rodando na porta localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Manipulador para a página principal
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, nil)
}

// Manipulador para salvar os dados no formato CHAVE:CONTEÚDO
func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		key := r.FormValue("key")
		value := r.FormValue("value")
		data := fmt.Sprintf("%s:%s\n", key, value)
		file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, "Unable to save data", http.StatusInternalServerError)
			return
		}
		defer file.Close()
		if _, err := file.WriteString(data); err != nil {
			http.Error(w, "Unable to save data", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Data saved successfully")
	}
}

// Manipulador para recuperar os dados com base na chave fornecida
func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		key := r.URL.Query().Get("key")
		content, err := os.ReadFile("data.txt")
		if err != nil {
			http.Error(w, "Unable to read data", http.StatusInternalServerError)
			return
		}
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 && parts[0] == key {
				fmt.Fprintf(w, "Content: %s", parts[1])
				return
			}
		}
		http.Error(w, "Key not found", http.StatusNotFound)
	}
}
