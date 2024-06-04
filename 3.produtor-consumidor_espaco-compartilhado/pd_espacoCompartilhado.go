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

var wg sync.WaitGroup

const capacidadeGalpao = 10

type Produto struct {
	ID          int
	NomeProduto string
}

var buffer []Produto
var mutex = &sync.Mutex{}
var cond = sync.NewCond(mutex)

func main() {
	buffer = make([]Produto, 0, capacidadeGalpao) //iniciando o buffer

	numeroProdutores := 3
	numeroConsumidores := 2

	wg.Add(numeroProdutores + numeroConsumidores) //inicializando o waitgroup
	for i := 1; i <= numeroProdutores; i++ {
		go produtor(i)
	}
	for i := 1; i <= numeroConsumidores; i++ {
		go consumidor(i)
	}

	wg.Wait()
}

func produtor(id int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		produto := Produto{ //criando um produto
			ID:          id*100 + i,
			NomeProduto: fmt.Sprintf("produto : %d", id*100+1),
		}

		mutex.Lock() //bloqueia o buffer
		for len(buffer) == capacidadeGalpao {//verificar se o tamanho do buffer e igual capacidade no galpão
			cond.Wait()//toda vez que essa condição e atingida ele bloqueia o buffer5
		}
		buffer = append(buffer, produto) //o produto criado vai ao buffer
		fmt.Printf("Produtor %d adicionou: %s\n", id, produto.NomeProduto)

		cond.Signal()  //Avisa sobre itens novos no buffer
		mutex.Unlock() //Desbloqueia o acesso do buffer

		time.Sleep(time.Millisecond * 500) //simulação do tempo para criar um item
	}
}

func consumidor(id int) {//inicia a função do consumidor
	defer wg.Done()//avisa quando a ação do consumidor e finalizada
	for i := 0; i < 5; i++ {
		mutex.Lock()//utiliza o mutex para bloquear o buffer

		for len(buffer) == 0 {
			cond.Wait()
		}

		produto := buffer[0]
		buffer = buffer[1:]
		fmt.Printf("Consumidor %d retirou: %s\n", id, produto.NomeProduto)

		cond.Signal()

		mutex.Unlock()

		time.Sleep(time.Millisecond * 500)
	}
}
