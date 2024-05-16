// Uma outra forma de verificar se o canal esta fechado é usando  função "range":
package main

import "fmt"

func produzir(c chan int) {
	c <- 1
	c <- 2
	c <- 3

	close(c)
}

func main() {
	c := make(chan int, 3)
	go produzir(c)

	for valor := range c {
		fmt.Println(valor)
	}
}
