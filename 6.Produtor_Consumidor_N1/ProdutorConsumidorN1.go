//Luis Thiago Silva Rabello

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Estrutura para representar os itens
type Item struct {
	ID      int
	Message string
}

// Função Produtor
func Produtor(id int, ch chan<- Item, done <-chan bool) {
	for {
		// Gera um novo item
		item := Item{ID: id, Message: fmt.Sprintf("Mensagem do Produtor %d", id)}

		// Envia o item para o canal
		ch <- item

		// Simula um tempo de produção aleatório
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		select {
		case <-done:
			// Sinaliza que o produtor deve parar
			fmt.Printf("Produtor %d finalizado\n", id)
			return
		default:
			// Continua a produzir itens
		}
	}
}

// Função principal
func main() {
	// Define o número de produtores
	numProdutores := 4

	// Cria o canal de comunicação
	ch := make(chan Item, numProdutores)

	// Cria o canal de sinalização de término
	done := make(chan bool)

	// Inicia as goroutines dos produtores
	for i := 0; i < numProdutores; i++ {
		go Produtor(i+1, ch, done)
	}

	// Inicia a gorotina do consumidor
	go Consumidor(ch, done)

	// Espera um tempo para que os produtores e o consumidor concluam suas tarefas
	time.Sleep(time.Second * 5)

	// Sinaliza o término para os produtores e o consumidor
	close(done)

	// Aguarda a finalização das goroutines
	time.Sleep(time.Second)
}

// Função Consumidor
func Consumidor(ch <-chan Item, done <-chan bool) {
	for {
		// Recebe um item do canal
		item, ok := <-ch

		if !ok {
			// O canal foi fechado, sinalizando que todos os produtores finalizaram
			fmt.Println("Consumidor finalizado")
			return
		}

		// Processa o item recebido
		fmt.Printf("Consumidor recebeu: ID = %d, Mensagem = %s\n", item.ID, item.Message)
	}
}
