/*
Exercício: Produtor/Cosumidor(1-M): Usando Canal para Sincronizar
Neste trabalho deve ser implementada uma função Produtor  e uma função Consumidor, utilizando o esquema de  1 Produtor para M Consumidores, usando como mecanismo de sincronização a estrutura de Canal oferecida pela linguagem Go.

1) A função  Produtor deve funcionar como uma gorotina independente e que armazenar itens no canal

2) A função  Consumidor deve funcionar como uma gorotina independente e que retira itens do canal.
*/

//Luis Thiago Silva Rabello

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// capacidade do buffer do canal
const capacidadeBuffer = 10

// struct da mensagem
type Mensagem struct {
	ID       int
	Conteudo string
}

func main() {
	var wg sync.WaitGroup

	canal := make(chan Mensagem, capacidadeBuffer)

	// inicia a goroutine de produtor
	wg.Add(1)
	go produtor(canal, &wg)

	// número de consumidores
	numConsumidores := 3

	// inicia múltiplas goroutines de consumidores
	for i := 1; i <= numConsumidores; i++ {
		wg.Add(1)
		go consumidor(canal, &wg, i)
	}

	// Aguarda a conclusão das goroutines
	wg.Wait()
	fmt.Println("Fim da execução")
}

// Função Produtor que armazena mensagens no canal
func produtor(canal chan Mensagem, wg *sync.WaitGroup) {
	defer wg.Done() // Marca a conclusão da goroutine ao final

	for i := 1; i <= 20; i++ {
		time.Sleep(time.Second) // Simula o tempo de escrita da mensagem
		mensagem := Mensagem{   // Cria uma nova mensagem com um ID e um conteúdo
			ID:       i,
			Conteudo: fmt.Sprintf("Conteúdo da Mensagem %d", i),
		}
		canal <- mensagem // Coloca a mensagem no canal
		fmt.Printf("Produtor escreveu esta mensagem: %v\n", mensagem)
	}
	close(canal) // Fecha o canal após criar todas as mensagens
}

// Função Consumidor que retira e lê mensagens do canal
func consumidor(canal chan Mensagem, wg *sync.WaitGroup, id int) {
	defer wg.Done() // Marca a conclusão da goroutine ao final

	for mensagem := range canal {
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) // Simula o tempo de leitura
		fmt.Printf("Consumidor %d: Leu esta mensagem: %v\n", id, mensagem)
	}
}
