package main

import "fmt"

type Location struct {
	Lat float64 `json: lat`
	Long float64 `json: long`
}

func main() {
	loc := Location{30.54554, 23.4050}
	fmt.Println(loc)
}
