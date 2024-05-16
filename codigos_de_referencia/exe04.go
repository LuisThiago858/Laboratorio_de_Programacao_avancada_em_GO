// Imprimindo valor a mais do que o canal produziu:
package main

import "fmt"

func produzir(c chan int) {
	c <- 101
	c <- 226
	c <- 382

	//close(c)
}

func main() {
	c := make(chan int, 3)
	go produzir(c)

	fmt.Println(<-c, <-c, <-c, <-c)

}
