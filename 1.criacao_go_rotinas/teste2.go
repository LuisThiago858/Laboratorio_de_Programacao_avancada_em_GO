package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Gera uma matriz de inteiros aleatorios com dimensões x por y passado na main
func gerarMatriz(x, y int) [][]int {
	matriz := make([][]int, x)
	for i := range matriz {
		matriz[i] = make([]int, y)
		for j := range matriz[i] {
			matriz[i][j] = rand.Intn(100) // valores aleatorios entre 0 e 99
		}
	}
	return matriz
}

func multiplicaMatrizesParelelo(matriz1, matriz2 [][]int) [][]int {
	x := len(matriz1)
	y := len(matriz1[0])
	z := len(matriz2[0])

	matrizResultanteParalela := make([][]int, x)
	for i := range matrizResultanteParalela {
		matrizResultanteParalela[i] = make([]int, z)
	}
	var wg sync.WaitGroup //Waitgroup serve para esperar que uma coleção de goroutines termine sua execução.
	//funcao que calcula o valor de um elemento da matriz resultante
	calculaElemento := func(i, j int) {
		defer wg.Done()
		for k := 0; k < y; k++ {
			matrizResultanteParalela[i][j] += matriz1[i][k] * matriz2[k][j]
		}
	}

	//Lança uma goroutine para cada elemento da matriz resultante
	for i := 0; i < x; i++ {
		for j := 0; j < z; j++ {
			wg.Add(1)
			go calculaElemento(i, j)
		}
	}
	//Espera todas as gorutines terminarem
	wg.Wait()

	return matrizResultanteParalela
}

func main() {
	/*
		matriz1 := [][]int{
			{15, 23, 35, 42},
			{47, 58, 66, 67},
			{72, 80, 91, 93},
		}
		matriz2 := [][]int{
			{10, 11, 12, 15},
			{23, 44, 1},
			{16, 17, 18},
		}
	*/
	matriz1 := gerarMatriz(50, 50)
	matriz2 := gerarMatriz(50, 50)

	comecotParalelo := time.Now()
	CPar := multiplicaMatrizesParelelo(matriz1, matriz2)
	tempoPar := time.Since(comecotParalelo)

	fmt.Println("Matriz resultante paralela:")
	for _, linha2 := range CPar {
		fmt.Println(linha2)
	}
	fmt.Println("Tempo de execução da matriz resultante paralela foi de: ", tempoPar)
}
