package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU(), "Numero de Processadores") //numero de processadores
	fmt.Println(runtime.NumGoroutine(), " Goroutines antes") //numero de goroutines ou threads antes de chamar as funcoes, que e a main
	//add(total de funcoes)
	wg.Add(2) //programa tem uma goroutine que o programa tem que esperar

	go func1() //se houver o go na frente essa funcao e executada apenas
	go func2()

	fmt.Println(runtime.NumGoroutine(), " Goroutines depois") //numero de goroutines ou threads depois de chamar as funcoes, e termina com 3 contando a main
	//espera, antes do programa encerrar
	wg.Wait()
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1: ", i)
		time.Sleep(20)
	}
	//deu! Concluido
	wg.Done()
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2: ", i)
		time.Sleep(20)
	}
}
