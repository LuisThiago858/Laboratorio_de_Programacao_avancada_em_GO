/*
Exercício: Produtor/Cosumidor: Concorrencia usando Canais
Assuma que um produtor vai ao depósito armazenar as caixas que vai
 produzindo. Da mesma forma, um consumidor vai ao depósito para retirar
  caixas que vai consumir.

As seguintes atividades devem ser feitas:

1) Implemente uma função Produtor que coloque um item de dado no canal,
o canal inicialmente esta vazio; Não é possível armazenar nenhuma item
caso a capacidade máxima já tenha sido atingida. Da mesma forma,
não é possível retirar nenhum item, caso o canal esteja vazio.

2) A função  Produtor deve funcionar como uma gorotina independente e
que armazenar itens no canal.

3) A função  Consumidor deve funcionar como uma gorotina independente
e que retira itens do canal.

4) Implemente uma função extra para limpar o canal, removendo todos
os itens e fechar o mesmo.

5) Implemente uma função extra para preencher  todo o canal,
causando a suspensão da inserção. Ela deve ser usada para finalizar
a produção e fechar o canal.


OBSERVAÇÕES:
1) Pode pegar idéias e trechos de códigos da internet, mas tem que
comentar de onde copiou o trecho, além disso tem que acrerscentar novos
trechos de codigo para não configurar plágio.
2) Comente todas as funções no programa.
3) Enviar o codigo fonte e um exemplo da execução.
4) Coloque no codigo o nome dos membros da equipe.


*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// capacidade do galpão que basicamente o buffer do canal
const capacidadeGalpao = 10

// struct do produto
type Produto struct {
	ID          int
	NomeProduto string
}

func limparCanal(canal chan Produto) {
	for len(canal) > 0 {
		<-canal //remove produtos do canal ate que esvazie completamente
	}
	close(canal)
	fmt.Println("Canal limpo e fechado")
}

// Função para preencher o canal ate a capacidade maxiama
func preencherCanal(canal chan Produto) {
	for i := 1; 1 <= capacidadeGalpao; i++ {
		produto := Produto{ // Cria um novo produto com um ID e um nome
			ID:          i,
			NomeProduto: fmt.Sprintf("Conteudo da caixa %d", i),
		}
		canal <- produto //aloca esse produto no canal
		fmt.Printf("Preencheu o canal com o produto %v\n", produto)
	}
}
func main() {
	canal := make(chan Produto, capacidadeGalpao)

	//inicia a goroutine de produtor
	wg.Add(1)
	go produtor(canal, &wg)
	//inicia a gorountine de consumidor
	wg.Add(1)
	go consumidor(canal, &wg)
	//Aguarda um tempo para preencher o canal
	time.Sleep(15 * time.Second)
	//exemplo de preenchimento do canal
	preencherCanal(canal)
	//aguarda um tempo para limpar o canal
	time.Sleep(5 * time.Second)
	//Exemplo de limpeza do canal
	limparCanal(canal)
	//Aguarda um tempo para que todas as goroutines terminem
	time.Sleep(2 * time.Second)

	//fecha o canal apos limpar para evitar que seja fechada diversas vezes
	close(canal)
	fmt.Println("Canal fechado")

	//Aguarda a conclusão das goroutine
	wg.Wait()
	fmt.Println("Fim da execucao")
}

// Função Produtor que armazena produtos no canal
func produtor(canal chan Produto, wg *sync.WaitGroup) {
	defer wg.Done() //Marca a conclusão da goroutine ao final

	for i := 1; i <= 20; i++ {
		time.Sleep(time.Second) //Simula o tempo de produção
		produto := Produto{     //armazena os valores da struct em uma variavel
			ID:          i,
			NomeProduto: fmt.Sprintf("Conteudo do Produto %d", i),
		}
		canal <- produto //ALoca o produto no canal
		fmt.Printf("Produtor armazenou este produto %d\n", produto)
	}
	close(canal)
}

// Função consumidor que retira produtos do canal
func consumidor(canal chan Produto, wg *sync.WaitGroup) {

	defer wg.Done() //marca a conclusão da goroutine ao final
	for produto := range canal {
		time.Sleep(2 * time.Second) //Simula o tempo de consumo
		fmt.Printf("Consumidor: Retirou este produto %v\n", produto)
	}
}
