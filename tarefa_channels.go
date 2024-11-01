package main

import (
	"fmt"
	"math"
	"sync"
)

// Função para verificar a primalidade de um número
func ehPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n)))+1; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Função executada pelas goroutines para verificar a primalidade dos números recebidos pelo canal
func verificaPrimos(num_channel <-chan int, resultado chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() //é chamado por cada goroutine quando ela termina seu trabalho, decrementando o contador interno do grupo de espera.
	for num := range num_channel {
		if ehPrimo(num) {
			resultado <- 1 // Envia 1 para indicar que encontrou um primo
		}
	}
}

func main() {
	var N, M int
	fmt.Print("Digite a quantidade da sequencia: ")
	fmt.Scan(&N)
	fmt.Print("Digite o número de goroutines: ")
	fmt.Scan(&M)

	// Canais para enviar números e receber resultados
	num_channel := make(chan int, M)
	resultado := make(chan int, M)

	var wg sync.WaitGroup //é uma estrutura que permite sincronizar a execução de goroutines, garantindo que o fluxo principal aguarde a conclusão de todas as goroutines antes de continuar ou finalizar

	// Inicia M goroutines para verificar a primalidade dos números
	for i := 0; i < M; i++ {
		wg.Add(1) //é usado para adicionar o número de goroutines que precisam ser aguardadas.
		go verificaPrimos(num_channel, resultado, &wg)
	}

	// Envia números de 1 a N para o canal né usado para bloquear a execução da goroutine até que todas as goroutines tenham chamado Done().umChan
	go func() {
		for i := 1; i <= N; i++ {
			num_channel <- i
		}
		close(num_channel)
	}()

	// Lê os resultados e conta os primos
	go func() {
		wg.Wait() //é usado para bloquear a execução da goroutine até que todas as goroutines tenham chamado Done().
		close(resultado)
	}()

	// Contador total de números primos
	totalPrimos := 0
	for primo := range resultado {
		totalPrimos += primo
	}

	fmt.Printf("Total de números primos encontrados de 1 a %d: %d\n", N, totalPrimos)
}
