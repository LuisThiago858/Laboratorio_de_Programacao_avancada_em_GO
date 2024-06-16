/*
Exercício: Produtor/Cosumidor(1-1): Usando Canal para Sincronizar
Neste trabalho deve ser implementada uma função Produtor  e uma função Consumidor, utilizando o esquema de  1 Produtor para 1 Consumidor,
 usando como mecanismo de sincronização a estrutura de Canal oferecida pela linguagem Go.

1) A função  Produtor deve funcionar como uma gorotina independente e que armazenar itens no canal

2) A função  Consumidor deve funcionar como uma gorotina independente e que retira itens do canal.


OBSERVAÇÕES:
1) Pode pegar idéias e trechos de códigos da internet, mas tem que comentar de onde copiou o trecho, além disso tem que acrerscentar
 novos trechos de codigo para não configurar plágio.
2) Coloque Comentários em todas as funções no programa.
3) Enviar o codigo fonte com um exemplo da execução.
4) Coloque no codigo o nome de todos os membros da equipe.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// capacidade do galpao que basicamente é o buffer do canal
const capacidadeGalpao = 10

// struct do produto
type Produto struct {
	ID          int
	NomeProduto string
}

func main() {
	var wg sync.WaitGroup

	canal := make(chan Produto, capacidadeGalpao)

	// inicia a goroutine de produtor
	wg.Add(1)
	go produtor(canal, &wg)

	// inicia a goroutine de consumidor
	wg.Add(1)
	go consumidor(canal, &wg)

	// Aguarda a conclusão das goroutines
	wg.Wait()
	fmt.Println("Fim da execução, galpão esvaziado")
}

// Função Produtor que armazena produtos no canal
func produtor(canal chan Produto, wg *sync.WaitGroup) {
	defer wg.Done() // Marca a conclusão da goroutine ao final

	for i := 1; i <= 20; i++ {
		time.Sleep(time.Second) // Simula o tempo de produção
		produto := Produto{     // Cria um novo produto com um ID e um nome
			ID:          i,
			NomeProduto: fmt.Sprintf("Conteúdo do Produto %d", i),
		}
		canal <- produto // Coloca o produto no canal
		fmt.Printf("Produtor armazenou este produto: %v\n", produto)
	}
	close(canal) // Fecha o canal após produzir todos os itens
}

// Função Consumidor que retira produtos do canal
func consumidor(canal chan Produto, wg *sync.WaitGroup) {
	defer wg.Done() // Marca a conclusão da goroutine ao final

	for produto := range canal {
		time.Sleep(2 * time.Second) // Simula o tempo de consumo
		fmt.Printf("Consumidor: Retirou este produto: %v\n", produto)
	}
}
