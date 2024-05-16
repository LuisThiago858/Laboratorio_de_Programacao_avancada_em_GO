package main

import (
    "fmt"
    "sync"
)

func main() {
    // Lista de números
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Canal para enviar e receber resultados das goroutines
    resultCh := make(chan int)

    // WaitGroup para sincronizar as goroutines
    var wg sync.WaitGroup

    // Número de goroutines que serão criadas
    numGoroutines := 4

    // Dividir a lista em partes iguais para cada goroutine
    chunkSize := len(numbers) / numGoroutines

    // Loop para criar e iniciar as goroutines
    for i := 0; i < numGoroutines; i++ {
        start := i * chunkSize
        end := (i + 1) * chunkSize
        if end > len(numbers) {
            end = len(numbers)
        }

        // Incrementar WaitGroup
        wg.Add(1)

        // Goroutine para calcular a soma parcial
        go func(nums []int) {
            defer wg.Done()
            sum := 0
            for _, num := range nums {
                sum += num
            }
            // Enviar a soma parcial para o canal
            resultCh <- sum
        }(numbers[start:end])
    }

    // Goroutine para fechar o canal quando todas as goroutines terminarem
    go func() {
        wg.Wait()
        close(resultCh)
    }()

    // Soma final
    totalSum := 0

    // Loop para receber e somar os resultados das goroutines
    for partialSum := range resultCh {
        totalSum += partialSum
    }

    // Imprimir a soma total
    fmt.Println("Soma Total:", totalSum)
}
