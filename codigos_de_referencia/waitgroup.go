//Exemplo de ProdutorConsumidor usando Canais e Select:

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

/*
a variável wg é do tipo para realizara a sincronização dos processos
Um WaitGroup espera que uma coleção de goroutines termine.
A goroutine principal chama Add para definir o número de goroutines
a serem aguardadas. Em seguida, cada uma das goroutines é executada
e chama Done quando terminar. Ao mesmo tempo, Wait pode ser usado
para bloquear até que todas as goroutines terminem.
*/

func main() {

	wg.Add(2) // Defini quantos gorotinas serão aguardadas

	go func1()
	go func2()

	wg.Wait() //aguarda todas as rotinas enviarem um Done

}

func func1() {
	for i := 0; i < 15; i++ {
		fmt.Println("func1:", i)
	}
	wg.Done() //sinaliza que acabou a execução
}

func func2() {
	for i := 0; i < 15; i++ {
		fmt.Println("func2:", i)
	}
	wg.Done() //sinaliza que acabou a execução
}
