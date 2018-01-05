package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World\nDigite seu nome")
	var resp string
	fmt.Scanln(&resp)
	fmt.Println(resp)
}
