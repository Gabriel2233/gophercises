package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// Bitwise shifting any number x to the left by y bits yields x * 2 ** y

func main() {
	data := []int{23, 75, 34, 100}
	w, h := len(data)*60+1, 100 // 250

	rect := image.Rect(0, 0, w, h)
	img := image.NewRGBA(rect)

	grey := image.NewUniform(color.RGBA{240, 240, 240, 255})
	blue := image.NewUniform(color.RGBA{180, 180, 250, 255})

	draw.Draw(img, rect, grey, image.Point{0, 0}, draw.Src)

	for i, dp := range data {
		x0, y0 := (i*60 + 10), 100-dp
		x1, y1 := (i+1)*60-1, 100

		bar := image.Rect(x0, y0, x1, y1)

		draw.Draw(img, bar, blue, image.Point{0, 0}, draw.Src)
	}

	outFile, err := os.Create("bar.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		panic(err)
	}
}
