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
	// Define manipuladores para diferentes rotas
	http.HandleFunc("/", indexHandler)            // Manipulador para a página principal
	http.HandleFunc("/save", saveHandler)         // Manipulador para salvar os dados
	http.HandleFunc("/retrieve", retrieveHandler) // Manipulador para recuperar os dados

	// Mensagem para indicar que o servidor está rodando
	fmt.Println("Server está rodando na porta localhost:8080")

	// Inicia o servidor web na porta 8080
	http.ListenAndServe(":8080", nil)
}

// Manipulador para a página principal
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Parseia o arquivo HTML e o renderiza
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, nil)
}

// Manipulador para salvar os dados no formato CHAVE:CONTEÚDO
func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Obtém os valores da chave e conteúdo do formulário
		key := r.FormValue("key")
		value := r.FormValue("value")
		// Formata a string para salvar no arquivo
		data := fmt.Sprintf("%s:%s\n", key, value)
		// Abre o arquivo data.txt para anexar dados
		file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// Retorna um erro se não for possível abrir o arquivo
			http.Error(w, "Não habilitado para salvar", http.StatusInternalServerError)
			return
		}
		defer file.Close()
		// Escreve a string formatada no arquivo
		if _, err := file.WriteString(data); err != nil {
			// Retorna um erro se não for possível escrever no arquivo
			http.Error(w, "Não habilitado para salvar", http.StatusInternalServerError)
			return
		}
		// Confirma que os dados foram salvos com sucesso
		fmt.Fprintf(w, "Dados salvos com sucesso")
	}
}

// Manipulador para recuperar os dados com base na chave fornecida
func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Obtém a chave da query string
		key := r.URL.Query().Get("key")
		// Lê o conteúdo do arquivo data.txt
		content, err := os.ReadFile("data.txt")
		if err != nil {
			// Retorna um erro se não for possível ler o arquivo
			http.Error(w, "Não habilitado para leitura", http.StatusInternalServerError)
			return
		}
		// Divide o conteúdo do arquivo em linhas
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			// Divide cada linha no primeiro ':' encontrado
			parts := strings.SplitN(line, ":", 2)
			// Verifica se a chave corresponde e se a linha contém os dois componentes
			if len(parts) == 2 && parts[0] == key {
				// Retorna o conteúdo correspondente
				fmt.Fprintf(w, "Conteúdo: %s", parts[1])
				return
			}
		}
		// Retorna um erro se a chave não for encontrada
		http.Error(w, "Chave não encontrada", http.StatusNotFound)
	}
}
