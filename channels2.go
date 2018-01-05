package main 

import (
	"fmt"
	"time"
)

func sendMessage(msg chan string, content string) {
	time.Sleep(time.Second * 2)
	msg <- content
}

func receiveMessage(msg chan string, done chan bool){
	fmt.Println("Uma mensagem recebida: ", <-msg)
	done <- true
}

func main() {
	msg := make(chan string)
	done := make(chan bool)

	go sendMessage(msg, "The boss wants you learn golang...")
	go receiveMessage(msg, done)

	<-done
}