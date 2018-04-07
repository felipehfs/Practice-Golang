package main

import (
	"fmt"
	"image"
) 

func main() {
	// Cria uma imagem em branco com 10 pixels de largura e 4 pixels de altura
	myImg := image.NewRGBA(image.Rect(0, 0, 10, 4))

	// Você pode acessar o pixels através myImage.Pix[i]
	myImg.Pix[0] = 255
	myImg.Pix[1] = 0
	myImg.Pix[2] = 0
	myImg.Pix[3] = 255

	// myImg.Pix contêm todos os pixels
	// em um array de uma única dimensão
	fmt.Println(myImg.Pix)

	fmt.Println(myImg.Stride)	
}
