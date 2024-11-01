package main

import (
	"fmt"
)

func tarefa(str chan string) {
	// Recebe a mensagem da Main e imprime
	msg := <-str
	fmt.Println("Goroutine:", msg)

	// Envia a resposta para Main
	str <- "Oi Main, bom dia, tudo bem?"

	// Recebe a segunda mensagem da Main e imprime
	msg = <-str
	fmt.Println("Goroutine:", msg)

	// Envia a última resposta para Main
	str <- "Certo, entendido."

	// Imprime a mensagem de finalização da Goroutine
	fmt.Println("Goroutine: finalizando")
}

func main() {
	// Cria um canal de comunicação não-bufferizado
	str := make(chan string)

	// Cria uma goroutine que executará a função 'tarefa'
	go tarefa(str)

	// Envia a primeira mensagem para a Goroutine
	str <- "Olá, Goroutine, bom dia!"

	// Lê a primeira resposta da Goroutine
	msg := <-str
	fmt.Println("Main:", msg)

	// Envia a segunda mensagem para a Goroutine
	str <- "Tudo bem! Vou terminar tá?"

	// Lê a segunda resposta da Goroutine
	msg = <-str
	fmt.Println("Main:", msg)

	// Imprime a mensagem de finalização da Main
	fmt.Println("Main: finalizando")
}
