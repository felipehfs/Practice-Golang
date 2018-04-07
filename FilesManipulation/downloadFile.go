package main 

import (
	"os"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Função para a manipulação de erro
func errorHandle(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// DownloadContent baixa o conteúdo da internet e grava no disco
func DownloadContent(url, filename string, done chan bool) {
	newFile, err := os.Create(filename)
	errorHandle(err)
	defer newFile.Close()
	
	response, err := http.Get(url)
	defer response.Body.Close()
	errorHandle(err)

	numBytesWritten, err :=  io.Copy(newFile, response.Body)
	errorHandle(err)
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	done <- true
}

func main(){
	done := make(chan bool)
	url := os.Args[1]
	filename := os.Args[2]
	go DownloadContent(url, filename, done)
	for {
		select {
		case <-done:
			fmt.Println()
			return
		default:
			fmt.Printf("\rLoading...")
		}

	}
}
