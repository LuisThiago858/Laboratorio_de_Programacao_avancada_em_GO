// Chamando a função Imprimindo duas vezes sequencial.
package main

import (
	"fmt"
	"time"
)

func imprimindo(n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", n)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	imprimindo(2)

	imprimindo(3)
}
