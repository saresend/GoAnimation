package main

import (
	"image/color"

	"github.com/Samuel-Resendez/GoAnimation"
)

func main() {

	col := color.RGBA{0, 0, 255, 255}

	for i := 0; i < 100; i++ {
		canvas := GoAnimation.CreateCanvas(800, 800, false)
		for j := 0; j < i; j++ {
			circles := GoAnimation.CreateCircle(200+5*j, 200, 20, col)
			GoAnimation.Draw(canvas, circles)
		}

		GoAnimation.SaveImage(&(*canvas).Src, i)
	}

}
