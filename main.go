package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {

	inputImageFilename := os.Args[1]
	reader, err := os.Open(inputImageFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	inputImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	//TODO: implement gaussian filter (or approximation of it)
}
