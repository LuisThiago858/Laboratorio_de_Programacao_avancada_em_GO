package main

import (
	"fmt"
) // Precisamos do pacote fmt para imprimir para o stdout

/*
Estas duas linhas seguintes criam dois novos canais, que são
uma dos construtores de concorrencia da linguagem golang.
O primeiro canal é um canal booleano o segundo é
um canal inteiro. Podemos ler ou gravar dados nos canais.
Vamos ver como em breve.
*/
var done = make(chan bool)
var msgs = make(chan int)

/*
Na função principal, são geradas duas "go" rotinas. "Go" rotinas
são funções leves que serão executadas simultaneamente pela linguagem em tempo de execução.
Elas são funções regulares, que ao usar a palavra-chave "go" antes
da chamada de função, faz com que as duas funções sejam executadas simultaneamente.
No final a thread principal é bloqueada lendo o canal "done".

Assim que algo vem pelo canal, nós lemos, o descartamos e
o tópico principal continua a executar. Neste caso, nós saímos do
O programa.
*/
func main() {
	go produce()
	go consume()
	<-done
}

/*
O produto executa um laço 10 vezes, excrevendo um inteiro (0..10) por vez no canal msg.
A cada escrita ele é bloqueado pelo canal e fica esperando até que o consumidor faça uma leituras do outro lado do canal.
Quando for produzido tudo no canal do produtor, um booleano no canal "done" é enviado para para liberar a rotina principal, por meio do canal "done".
*/
func produce() {
	for i := 0; i < 10; i++ {
		msgs <- i
	}
	done <- true
}

/*
The consume go routine loops infinitely and reads
on the msgs channel. It will block until something
comes in the channel.
The syntax can be a little bit strange for people
coming from other languages.
':=' creates a variable assigning it the type of the value
coming on the right of the assignation. An int in this
case.
'<-' is the go way to read from a channel.
Once we have the msg (int) we dump it in the stdout.
*/
func consume() {
	for {
		msg := <-msgs
		fmt.Println(msg)
	}
}
