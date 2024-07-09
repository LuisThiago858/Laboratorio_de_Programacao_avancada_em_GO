//Luis Thiago Silva Rabello

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Estrutura para representar os pedidos do restaurante
type Pedido struct {
	ID     int
	Prato  string
	Mesa   int
	Pronto bool
}

// Função auxiliar para gerar pratos aleatórios
func gerarPrato() string {
	pratos := []string{"Feijoada", "Picanha", "Strogonoff", "Pizza", "Lasanha"}
	return pratos[rand.Intn(len(pratos))]
}

// Função principal
func main() {
	// Define o número de cozinheiros e clientes
	numCozinheiros := 3
	numClientes := 5

	// Cria os canais de comunicação
	balcao := make(chan Pedido, numCozinheiros)
	done := make(chan bool)

	// Inicia as goroutines dos cozinheiros
	for i := 0; i < numCozinheiros; i++ {
		go Cozinheiro(i+1, balcao, done)
	}

	// Inicia as goroutines dos clientes
	for i := 0; i < numClientes; i++ {
		go Cliente(i+1, balcao, done)
	}

	// Espera um tempo para que os cozinheiros e clientes concluam suas tarefas
	time.Sleep(time.Second * 10)

	// Sinaliza o término para os cozinheiros e clientes fechando o canal `done`
	close(done)

	// Aguarda a finalização das goroutines
	time.Sleep(time.Second)
}

// Função Produtor (Cozinheiro)
func Cozinheiro(id int, ch chan Pedido, done <-chan bool) {
	for {
		// Gera um novo pedido
		pedido := Pedido{ID: id, Prato: gerarPrato(), Mesa: rand.Intn(10) + 1}

		// Simula tempo de preparo do pedido
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)

		// Marca o pedido como pronto
		pedido.Pronto = true

		// Envia o pedido para o balcão
		ch <- pedido

		// Verifica se deve parar de cozinhar
		select {
		case <-done:
			fmt.Printf("Cozinheiro %d finalizado\n", id)
			return
		default:
			// Continua cozinhando
		}
	}
}

// Função Consumidor (Cliente)
func Cliente(id int, ch <-chan Pedido, done <-chan bool) {
	for {
		// Recebe um pedido do balcão
		pedido, ok := <-ch

		if !ok {
			// O balcão foi fechado, sinalizando que todos os pedidos foram entregues
			fmt.Printf("Cliente %d finalizado\n", id)
			return
		}

		// Verifica se o pedido está pronto
		if pedido.Pronto {
			fmt.Printf("Cliente %d consumindo pedido %d: %s (Mesa %d)\n", id, pedido.ID, pedido.Prato, pedido.Mesa)
		} else {
			fmt.Printf("Cliente %d esperando pedido %d: %s (Mesa %d)\n", id, pedido.ID, pedido.Prato, pedido.Mesa)
			time.Sleep(time.Second) // Simula tempo de espera do cliente
		}
	}
}
