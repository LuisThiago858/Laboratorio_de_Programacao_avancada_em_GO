/*
Exercício: Produtor/Cosumidor: Usando Espaço Compartilhado
Em programas multithread, muitas vezes há uma divisão de trabalho entre threads. Em um padrão comum, alguns threads são produtores
e outros são consumidores. Os produtores criam algum tipo de item e os adicionam a uma estrutura de dados; os consumidores removem
 os itens e os processam.
Os programas orientados a eventos são um bom exemplo. Um “evento” é algo que acontece e exige que o programa
 responda: o usuário pressiona uma tecla ou move o mouse, um bloco de dados chega do disco, um pacote chega da rede,
  uma operação pendente é concluída.
Sempre que ocorre um evento, um thread produtor cria um objeto de evento e o adiciona ao buffer de eventos. Ao mesmo tempo,
 os threads consumidores retiram eventos do buffer e os processam. Nesse caso, os consumidores são chamados de “manipuladores
  de eventos”.
Existem várias restrições de sincronização que precisamos impor para que este sistema funcione corretamente:
- Enquanto um item está sendo adicionado ou removido do buffer, o buffer fica em um estado inconsistente. Portanto, os threads
 devem ter acesso exclusivo ao buffer.

- Se um thread consumidor chegar enquanto o buffer estiver vazio, ele será bloqueado até que um
produtor adiciona um novo item.

As seguintes atividades devem ser feitas:

1) Implemente uma função Produtor que coloque um item de dado espaço compartilhado, o qual esta inicialmente vazio; Não é possível
 armazenar nenhuma item caso a capacidade máxima já tenha sido atingida. Da mesma forma, não é possível retirar nenhum item,
  caso o espaco esteja vazio.

2) A função  Produtor deve funcionar como uma gorotina independente e que armazenar itens no espaço compartilhado.

3) A função  Consumidor deve funcionar como uma gorotina independente e que retira itens do espaço compartilhado.


OBSERVAÇÕES:
1) Pode pegar idéias e trechos de códigos da internet, mas tem que comentar de onde copiou o trecho, além disso tem que acrerscentar
 novos trechos de codigo para não configurar plágio.
2) Comente todas as funções no programa.
3) Enviar o codigo fonte e um exemplo da execução.
4) Coloque no codigo o nome dos membros da equipe.
*/

//Luis Thiago Silva Rabello

package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup para esperar a conclusão dos produtores e consumidores
var wg sync.WaitGroup

const capacidadeGalpao = 10 //capacidade maxima do buffer

type Produto struct { //Estrutura de um produto
	ID          int
	NomeProduto string
}

var buffer []Produto           //buffer compartilhado
var mutex = &sync.Mutex{}      //Mutex para proteger o acesso ao buffer
var cond = sync.NewCond(mutex) //Sincronização entre produtores e consumidores

// Main
func main() {
	buffer = make([]Produto, 0, capacidadeGalpao) //iniciando o buffer com o tamanho do galpão

	numeroProdutores := 3   //numero de produtores
	numeroConsumidores := 2 //numero de consumidores

	wg.Add(numeroProdutores + numeroConsumidores) //inicializando o waitgroup com o numero de totak de goroutines
	//Iniciando a goroutine de produtores
	for i := 1; i <= numeroProdutores; i++ {
		go produtor(i)
	}
	//Iniciamos a goroutine de consumidores
	for i := 1; i <= numeroConsumidores; i++ {
		go consumidor(i)
	}

	wg.Wait() //espera até que todas as goroutines estejam terminadas
}

// função produto
func produtor(id int) {
	defer wg.Done() //diz que essa será a ultima função que sera executada avisando que ela foi concluida
	for i := 0; i < 5; i++ {
		produto := Produto{ //criando um produto novo
			ID:          id*100 + i,                            //id multiplicado por cem  mais o i do for
			NomeProduto: fmt.Sprintf("produto : %d", id*100+1), //o nome do produto e o seu id
		}

		mutex.Lock()                          //bloqueia o acesso ao buffer
		for len(buffer) == capacidadeGalpao { //espera até ter espaco disponivel no buffer
			cond.Wait()
		}
		buffer = append(buffer, produto) //o produto e adicionado ao buffer
		fmt.Printf("Produtor %d adicionou: %s\n", id, produto.NomeProduto)

		cond.Signal()  //Avisa sobre itens novos no buffer
		mutex.Unlock() //Desbloqueia o acesso do buffer

		time.Sleep(time.Millisecond * 500) //simulação do tempo para criar um item
	}
}

// função consumidor
func consumidor(id int) {
	defer wg.Done() //diz que essa será a ultima função que sera executada avisando que ela foi concluida
	for i := 0; i < 5; i++ {
		mutex.Lock() //Bloqueia o acesso ao buffer

		//espera que tenha ao menos um item no buffer
		for len(buffer) == 0 {
			cond.Wait()
		}

		//retira o produto do buffer, decrementa o buffer e printa o quem retirou e o que retirou
		produto := buffer[0]
		buffer = buffer[1:]
		fmt.Printf("Consumidor %d retirou: %s\n", id, produto.NomeProduto)

		cond.Signal() //sinaliza que tem espaco no buffer

		mutex.Unlock() //bloqueia o acesso ao buffer

		time.Sleep(time.Millisecond * 500) //simula o tempo para processar a retirada do item
	}
}
