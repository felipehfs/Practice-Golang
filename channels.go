package main

import (
	"time"
	"fmt"
)
func task(done chan bool) {
	fmt.Println("Start task")
	time.Sleep(time.Second * 4)
	fmt.Println("End task")
	done <- true
}

func main() {
	done := make(chan bool)

	go task(done)

	<-done
}
