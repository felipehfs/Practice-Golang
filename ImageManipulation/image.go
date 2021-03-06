package main
import (
	"image"
	"image/color"
	"os"
	"image/png"
)

func main() {
	// Create an 100 x 50 image
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	// Draw a red dot at (2, 3)
	for y := 3; i < 36; y++{
		img.Set(2, y, color.RGBA{255, 0, 0, 255})
	}
	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
