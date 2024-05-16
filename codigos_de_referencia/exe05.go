// Uma forma de evitar o deadlock é verificar se o canal foi fechado
package main

import "fmt"

func produzir(c chan int) {
	c <- 101
	c <- 226
	c <- 382
	close(c)
}

func main() {
	c := make(chan int, 3)
	go produzir(c)
	for {
		valor, ok := <-c
		if ok {
			fmt.Println(valor)
		} else {
			break
		}

	}

}
