/*
Exercício: Criação de Gorotinas
Crie duas aplicação implementadas usando a linguagem Go para um mesmo
problema, sendo que as implementações serão uma sequencial e outra
 paralela.

1. Escolha uma aplicação que resolva qualquer problema, que possa ter
uma opção de solução sequencial e uma opção paralela. Por exemplo
soma dos elementos de um vetor de 10.000 posições.

2. Faça a implementação sequencial e contabilize o tempo de execução
da solução no início e no final da execução.

3. Faça a implementação paralela e contabilize o tempo de execução da
solução no início e no final da execução.

4. Compare os tempos e comente sobre as soluções.

OBSERVAÇÕES:
1) Pode pegar idéias e trechos de códigos da internet, mas tem que
comentar de onde copiou o trecho, além disso tem que acrerscentar
novos trechos de codigo para não configurar plágio.
2) Comente todas as funções no programa.
3) Enviar o codigo fonte e um exemplo da execução.
4) Coloque no codigo o nome dos membros da equipe
*/

//Membros da equipe: Luis Thiago Silva Rabello

// multiplicação de matrizes Sequencial e Paralela
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
		defer wg.Done() //adia a execução deste trecho do codigo
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

	//medir, imprimir e mostrar o tempo de execução da matriz resultante sequencial
	comecoSequencial := time.Now()                         //pega o momento exato do comeco da execução do codigo
	RSeq := multiplicaMatrizesSequencial(matriz1, matriz2) //atribui o resultado da multiplicacao a essa variavel que e uma matriz
	tempoSeq := time.Since(comecoSequencial)

	fmt.Println("Matriz resultante sequencial:")
	for _, linha := range RSeq { //varre a matriz resultante
		fmt.Println(linha) //imprime o que estiver la dentro
	}
	fmt.Println("Tempo de execução da matriz resultante sequencial foi de: ", tempoSeq) //imprime o tempo corrido em milissegundos

	//medir, imprimir e mostrar o tempo de execução da matriz resultanete paralela

	comecotParalelo := time.Now()
	go multiplicaMatrizesParelelo(matriz1, matriz2)
	CPar := multiplicaMatrizesParelelo(matriz1, matriz2)
	tempoPar := time.Since(comecotParalelo)

	fmt.Println("Matriz resultante paralela:")
	for _, linha2 := range CPar {
		fmt.Println(linha2)
	}
	fmt.Println("Tempo de execução da matriz resultante paralela foi de: ", tempoPar)
}

/*neste caso a solução sequencial se mostrou muito mais eficiente do que a versão paralela para o calculo de matrizes
por causa da dependencia de dados e a carga de trabalho desbalanceada, respectivamente por essas operacoes não terem muitas operacoes
independentes que podem ser executadas simultaneamente assim sendo executadas de em uma ordem especifica. Outro motivo e que a thread pode
está terminando de maneira mais rápida e isso pode está deixando ela ociosa esperando outras terminarem e consequentemente reduzindo a eficiencia do processo
*/
