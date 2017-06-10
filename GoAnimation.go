package GoAnimation

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"path/filepath"
)

func init() {
	fmt.Println("Initializing package...")

	//Checks if Frames directory exists, and builds it if not
	filePath := filepath.Join(".", "Frames")
	os.MkdirAll(filePath, os.ModePerm)
}

//CreateCircle generates a circle Sketch
func CreateCircle(centerX, centerY, radius int, col color.RGBA) *Sketch {
	θ, θx := 0.0, 0.01
	var yDel, xDel int
	var pixels []Pixel
	for ; θ < 2.0*3.14159; θ += θx {
		for i := 0; i < radius; i++ {
			yDel, xDel = int(float64(i)*math.Sin(θ)), int(float64(i)*math.Cos(θ))
			newPixel := Pixel{xCoord: centerX + xDel, yCoord: centerY + radius - yDel, color: col}
			pixels = append(pixels, newPixel)
		}

	}
	return &Sketch{
		pixels: pixels,
	}
}

//CreateLine generates a line sketch
/*func CreateLine(startX, startY, endX, endY int) *Sketch {

}*/
