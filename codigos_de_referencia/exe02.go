// Chamando a função dormir mais lenta que a main.
package main

import (
	"fmt"
	"time"
)

func dormir(t int) {
	fmt.Println("Goroutine dormindo por", t, " segundos...")
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("Goroutine finalizada.")
}

func main() {
	go dormir(5)

	fmt.Println("Main dormindo por 3 segundos...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main finalizada.")
}
