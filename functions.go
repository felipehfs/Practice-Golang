package main

import (
	"fmt"
)

func media(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	total := media(1, 2, 5)
	fmt.Println(total)
}
