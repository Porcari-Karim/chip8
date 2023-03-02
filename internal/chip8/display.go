package chip8

import (
	//"github.com/faiface/pixel"
	//"github.com/faiface/pixel/imdraw"
	"image"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	//"image/png"
	//"os"
	//"fmt"
)
type Display struct {
	pixels []byte //
}
const (
    displayWidth  = 64 // Largeur du display en pixels
    displayHeight = 32 // Hauteur du display en pixels
)

func NewDisplay() *Display {
	return &Display{make([]byte, (32 * 64) / 8)}
}

func createPixels(display []byte) [][]color.RGBA {
	pixels := make([][]color.RGBA, displayHeight)
	for y := range pixels {
		pixels[y] = make([]color.RGBA, displayWidth)
	}

	for y := 0; y < displayHeight; y++ {
		for x := 0; x < displayWidth; x++ {
			byteIndex := (y * 8) + (x / 8)
			bitIndex := uint(x % 8)
			bitValue := (display[byteIndex] >> (7 - bitIndex)) & 1
			if bitValue == 0 {
				//fmt.Println("Generated black pixel: ", x, y, bitIndex, bitValue, display[byteIndex])
				pixels[y][x] = color.RGBA{0, 0, 0, 255}
			} else {
				//fmt.Println("Generated white pixel: ", x, y, bitIndex, bitValue, display[byteIndex])
				pixels[y][x] = color.RGBA{255, 255, 255, 255}
			}
		}
	}

	return pixels
}

func (e *Emulator) Draw(win *pixelgl.Window) {
	pixels :=  createPixels(e.display.pixels)// Tableau de pixels de couleurs

	// Créer une image à partir du tableau de pixels de couleurs
	img := image.NewRGBA(image.Rect(0, 0, 64, 32))
	for x := range pixels {
		for y:= 0; y < 64; y++ {
			img.Set(y, x, pixels[x][y])
		}
	}
	newImg := image.NewRGBA(image.Rect(0, 0, 640, 320))
	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			// trouver le pixel correspondant dans l'image originale
			srcColor := img.At(x, y)

			// copier la couleur dans la nouvelle image, en la répétant sur une grille de 10x10 pixels
			for i := 0; i < 10; i++ {
				for j := 0; j < 10; j++ {
					newX := x*10 + i
					newY := y*10 + j
					newImg.Set(newX, newY, srcColor)
				}
			}
		}
	}

	pic := pixel.PictureDataFromImage(newImg)
	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))


	// sauvegarder l'image dans un fichier PNG
	//file, err := os.Create("resized.png")
	//if err != nil {
	//	return
	//}
	//defer file.Close()
	//png.Encode(file, newImg)
	//r_file, err := os.Create("base.png")
	//if err != nil {
	//	return
	//}
	//defer r_file.Close()
	//png.Encode(r_file, img)

}
