package main 

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func main() {
	done := make(chan bool)
	
	go haskTask(done)
	wg.Add(1)
	go loader(done)

	wg.Wait()
}
func haskTask(done chan<- bool) {
	fmt.Println("Start task")
	time.Sleep(time.Second * 1)
	done <- true	
}

func loader(done <-chan bool) {
	defer wg.Done()

	i := 0
	load := []rune(`|\-/|`)
	for {
		select {
		case <-done:
			fmt.Printf("\r")
			fmt.Println("Done")
			return
		default:
			fmt.Printf("\r")
			fmt.Printf(string(load[i]))

			i++
			if i == len(load){
				i = 0
			}
		}
	}
}