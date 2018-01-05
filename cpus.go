package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Essa maquina possui %d core(s)\n", runtime.NumCPU())
}
