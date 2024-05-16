// Canal: o canal serve como meio de sincronização entre processos:
package main

import "fmt"

func produzir(c chan int) {
	c <- 48
}

func main() {
	c := make(chan int, 3)
	go produzir(c)

	valor := <-c //Atenção!!!!
	fmt.Println(valor)
}
