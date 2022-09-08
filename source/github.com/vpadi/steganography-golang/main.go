package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"strings"

	"github.com/vpadi/steganography-golang/steganography"
)

func readMessage() {

}

func writeMessage() {
	var pathToFile string
	var messageToWrite string

	fmt.Printf("Specify the path to the image file: ")
	fmt.Scanf("%s", &pathToFile)

	imageData := steganography.ReadImage(pathToFile)

	fmt.Printf("Specify the message to save in the image: ")
	fmt.Scanln(&messageToWrite)

	totalPixels := imageData.Bounds().Dx() * imageData.Bounds().Dy()

	if totalPixels*3 < len(messageToWrite)*8 {
		log.Fatal("Image size is smaller than the message to write")
	}

	imageBounds := imageData.Bounds()
	i := 0

	cypherImage := image.NewRGBA(image.Rect(0, 0, imageBounds.Dx(), imageBounds.Dy()))
	draw.Draw(cypherImage, cypherImage.Bounds(), imageData, imageBounds.Min, draw.Src)

	for y := imageBounds.Min.Y; y < imageBounds.Max.Y && i <= len(messageToWrite); y++ {
		for x := imageBounds.Min.X; x < imageBounds.Max.X && i <= len(messageToWrite); x++ {
			r, g, b, a := cypherImage.At(x, y).RGBA()

			letterBinary := byte(0)

			if i < len(messageToWrite) {
				letterBinary = messageToWrite[i]
			}

			r = r & uint32((letterBinary>>6)|0b11111100)
			g = g & uint32((letterBinary>>4)|0b11111100)
			b = b & uint32((letterBinary>>2)|0b11111100)
			a = a & uint32(letterBinary|0b11111100)

			cypherImage.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
			i++
		}
	}

	/* i = 0
	for y := imageBounds.Min.Y; y < imageBounds.Max.Y && i < len(messageToWrite); y++ {
		for x := imageBounds.Min.X; x < imageBounds.Max.X && i < len(messageToWrite); x++ {
			fmt.Println(cypherImage.At(x, y))
			i++
		}
	} */
}

func main() {
	var programMode string
	fmt.Print("Do you want to write or read a image message? ")
	fmt.Scanf("%s", &programMode)

	programMode = strings.ToLower(programMode)
	for programMode != "write" && programMode != "read" {
		fmt.Print("Please, enter write or read. ")
		fmt.Scanf("%s", &programMode)

		programMode = strings.ToLower(programMode)
	}

	if programMode == "read" {
		readMessage()
	} else {
		writeMessage()
	}
}
