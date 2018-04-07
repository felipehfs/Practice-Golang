package main

import (
	"fmt"
	"math"
)

func progressaoAritmetica(a1, r, n int) int {
	return a1 + (n - 1) * r
}

func jurosCompostos(c, i, n float64) float64 {
	return c * math.Pow((1 + i), n)
}

func main(){
	montante := jurosCompostos(2000.00, 0.043, 10)
	fmt.Println(montante)
}
