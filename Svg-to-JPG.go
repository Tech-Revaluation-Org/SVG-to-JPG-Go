package main

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	// Open the SVG file
	svgFile, err := os.Open("input.svg")
	if err != nil {
		panic(err)
	}
	defer svgFile.Close()

	// Decode the SVG image
	svgImg, _, err := image.Decode(svgFile)
	if err != nil {
		panic(err)
	}

	// Create a new blank raster image to render the SVG onto
	width := 800
	height := 600
	dest := image.NewRGBA(image.Rect(0, 0, width, height))

	// Draw the SVG onto the raster image
	gc := draw2dimg.NewGraphicContext(dest)
	draw2dimg.DrawSvg(gc, svgImg)

	// Create a new JPEG file
	jpegFile, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer jpegFile.Close()

	// Encode the raster image as JPEG and write it to the file
	err = jpeg.Encode(jpegFile, dest, &jpeg.Options{Quality: 100})
	if err != nil {
		panic(err)
	}

	fmt.Println("JPEG image generated successfully.")
}
