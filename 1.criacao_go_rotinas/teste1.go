package main

import (
	"fmt"
	"math/rand"
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

// funcao que que recebe duas matrizes, matriz1 e matriz2 e retornar uma matrizResultante
func multiplicaMatrizesSequencial(matriz1, matriz2 [][]int) [][]int {
	x := len(matriz1)
	y := len(matriz1[0])
	z := len(matriz2[0])

	//criacao da matriz3 resultante que servira de recipiente para matriz1 x matriz2
	matrizResultanteSequencial := make([][]int, x)
	for i := range matrizResultanteSequencial {
		matrizResultanteSequencial[i] = make([]int, z)
	}

	//multiplica as matrizes e armazena o resultado na matriz resultante
	for i := 0; i < x; i++ {
		for j := 0; j < z; j++ {
			for k := 0; k < y; k++ {
				matrizResultanteSequencial[i][j] += matriz1[i][k] * matriz2[k][j]
			}
		}
	}
	return matrizResultanteSequencial //retornar a matriz resultante
}

func main() {

	matriz1 := gerarMatriz(50, 50)
	matriz2 := gerarMatriz(50, 50)

	//medir, imprimir e mostrar o tempo de execução da matriz resultante sequencial
	comecoSequencial := time.Now()                         //pega o momento exato do comeco da execução do codigo
	RSeq := multiplicaMatrizesSequencial(matriz1, matriz2) //atribui o resultado da multiplicacao a essa variavel que e uma matriz
	tempoSeq := time.Since(comecoSequencial)

	fmt.Println("Matriz resultante sequencial:")
	for _, linha := range RSeq { //varre a matriz resultante
		fmt.Println(linha) //imprime o que estiver la dentro
	}
	fmt.Println("Tempo de execução da matriz resultante sequencial foi de: ", tempoSeq) //imprime o tempo corrido em milissegundos

}
