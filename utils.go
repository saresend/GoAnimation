package GoAnimation

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
)

//Stores Defaults for where frames are stored, and info about their format
const (
	Frames      = "Frames/"
	FrameFormat = ".png"
	EndLocation = "output.gif"
)

//Canvas is a wrapper for image.RGBA, that also stores what is drawn on it
type Canvas struct {
	Src      image.RGBA //The actual image itself
	Sketches []Sketch   //Some of the sketches (about the layers, it will evaluate them first to last)
	xDim     int        //Stores the dimensions for our canvas
	yDim     int
}

//Pixel A small struct for storing pixel info
type Pixel struct {
	xCoord int
	yCoord int
	color  color.RGBA
}

//A Sketch is used to represent a single image or visual element
type Sketch struct {
	pixels []Pixel //The pixels that compose a sketch
}

/*
 * Gets the url at which to store a frame, given its index
 */

func getFrameURL(number int) string {
	return Frames + strconv.Itoa(number) + FrameFormat
}

/*
 * Notes: Encodes the values as PNGs!
 * //TODO: Make it potentially variable
 */

//SaveImage is used for saving a raw image, with its index
func SaveImage(img *image.RGBA, number int) {

	f, err := os.Create(getFrameURL(number))
	defer f.Close() //Ensures that the file closes
	if err != nil {
		fmt.Println(err)
	}
	png.Encode(f, img)
}

/*
 * @Params: dimX, dimY, the dimensions of the canvas
 * @Params: isTransparent, will return a transparent canvas if true, and white if false
 *
 */

//CreateCanvas creates a blank canvas, set either transparent or white
func CreateCanvas(dimX int, dimY int, isTransparent bool) *Canvas {
	img := image.NewRGBA(image.Rect(0, 0, dimX, dimY))
	if !isTransparent {
		white := color.RGBA{255, 255, 255, 255}
		setFill(img, white, dimX, dimY)
	}
	return &Canvas{Src: *img, xDim: dimX, yDim: dimY}
}

/*
 * @Params: image.RGBA pointer, a color, and X and Y Bounds
 * Operates on an image.RGBA, rather than a Canvas
 */

func setFill(src *image.RGBA, col color.RGBA, boundX int, boundY int) {
	for i := 0; i < boundX; i++ {
		for j := 0; j < boundY; j++ {
			src.Set(i, j, col)
		}
	}
}

//Draw is used to add a sketch to a canvas
func Draw(canvas *Canvas, sketch *Sketch) {
	for i := 0; i < len(sketch.pixels); i++ {
		pixel := sketch.pixels[i]
		canvas.Src.Set(pixel.xCoord, pixel.yCoord, pixel.color)
	}
	canvas.Sketches = append(canvas.Sketches, *sketch)
}

func convertToPaletted(img *image.Image) *image.Paletted {
	pm, ok := img.(*image.Paletted)
}

//CompileGIF takes all the frames in the directory and generates the output gif
func CompileGIF() {
	files, err := ioutil.ReadDir(Frames)

	if err != nil {
		fmt.Println(err)
	}

	var frames []*image.Image
	var delays []int
	numFrames := len(files)
	for i := 0; i < numFrames; i++ {
		fileReader, _ := os.Open(files[i].Name())

		img, _ := png.Decode(fileReader)
		frames = append(frames, &img)
		delays = append(delays, smoothAnimation(i, numFrames))
	}

	f, _ := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: frames,
		Delay: delays,
	})
}

func smoothAnimation(frameNumber, maxFrame int) int {
	diff := maxFrame - frameNumber
	return int(1.0 / float64(diff))
}
