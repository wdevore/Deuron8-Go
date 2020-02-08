package imageloads

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	_ "image/png" // Register PNG format
	"log"
	"os"
	"testing"

	"github.com/fogleman/gg"
)

func TestRuns(t *testing.T) {
	// runGG(t)
	runLoad(t)
}

func runLoad(t *testing.T) {
	// infile, err := os.Open("testdata/test_rgb.png")
	infile, err := os.Open("../simulation/stimulus/letter_A.png")
	if err != nil {
		// replace this with real error handling
		log.Fatalln(err)
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	src, _, err := image.Decode(infile)
	if err != nil {
		// replace this with real error handling
		log.Fatalln(err)
	}

	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	fmt.Println("w: ", w, ", h: ", h)
	// .At(col, row)
	fmt.Println(src.At(0, 0))
	fmt.Println(src.At(1, 0))
	fmt.Println(src.At(2, 0))
	fmt.Println(src.At(3, 0))
	fmt.Println(src.At(4, 0))
}

func runGenImage(t *testing.T) {
	bounds := image.Rect(0, 0, 100, 100)
	rgb := image.NewRGBA(bounds)

	black := color.RGBA{0, 0, 0, 255}

	draw.Draw(rgb, rgb.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	row := 0
	for col := 0; col < bounds.Max.X; col++ {
		rgb.Set(col, row, color.White)
	}

	outfile, err := os.Create("testdata/out_rgb.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer outfile.Close()

	png.Encode(outfile, rgb)
}

func runGG(t *testing.T) {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("testdata/gg.png")
}
