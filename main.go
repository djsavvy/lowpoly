package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {

	inputImageFilename := flag.Arg(0)
	reader, err := os.Open(inputImageFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	inputImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inputImage)

	//TODO: implement gaussian filter (or approximation of it)
}
