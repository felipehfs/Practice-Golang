package main

import(
	"fmt"
	"log"
	"os"
)

func main() {
	// O Stat retorna as informações do arquivo 
	// se o arquivo não existir irá retornar um
	// erro.

	fileInfo, err := os.Stat("texto.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nome:", fileInfo.Name())
	fmt.Println("Tamanho:", fileInfo.Size())
	fmt.Println("Permissões:", fileInfo.Mode())
	fmt.Println("Ultimas modificações:", fileInfo.ModTime())
	fmt.Println("É um diretório:", fileInfo.IsDir())
	fmt.Printf("Tipo de interface de sistema %T\n", fileInfo.Sys())
	fmt.Printf("Informações do sistema: %+v\n\n", fileInfo.Sys())
}
