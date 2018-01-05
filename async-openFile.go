package main

import(
	"io/ioutil"
	"fmt"
	"os"
)

func ReadFile(file string, content chan string){
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	content <- string(data)
}

func main(){
	c1 := make(chan string)
	file1 := os.Args[1]
	go ReadFile(file1, c1)
	fmt.Println("Abrindo arquivos")

	fmt.Println(<-c1)
}
