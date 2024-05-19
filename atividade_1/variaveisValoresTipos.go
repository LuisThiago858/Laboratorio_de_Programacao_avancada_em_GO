package main

import (
	"fmt"
)

var y = "Olá Bom Dia" //variavel global

func main() {
	x := 10 //declaracao

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	x = 20 //atribuicao
	fmt.Printf("x: %v, %T\n", x, x)

}
