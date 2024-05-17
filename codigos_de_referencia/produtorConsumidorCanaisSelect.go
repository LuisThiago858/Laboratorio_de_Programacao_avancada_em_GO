package main

import "fmt"

/*
A função make aloca e inicializa um objeto do tipo slice, map ou chan (somente).
O primeiro argumento é um tipo, não um valor. Diferente de new, o tipo de retorno
de make é igual ao tipo de seu argumento, não um ponteiro para ele.
A especificação do resultado depende do tipo.
*/
var item = make(chan int)
var pronto = make(chan bool)
var fim = false

func produce() {
	for i := 0; i < 15; i++ {
		item <- i
	}
	pronto <- true
}

func consume(msg int) {
	fmt.Println(msg)
}

func main() {

	go produce()

	for !fim {
		select {
		case n := <-item:
			consume(n)
		case fim = <-pronto:
		}
	}

	fmt.Printf("Fim! \n")
}
