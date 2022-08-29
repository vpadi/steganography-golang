package steganography

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func ReadImage(pathToFile string) image.Image {
	reader, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)

	return m
}

func WriteImage(pathToFile string, imageData image.Image) {
	f, err := os.Create("outimage.png")
	if err != nil {
		// Handle error
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, imageData)
	if err != nil {
		log.Fatal(err)
	}
}
