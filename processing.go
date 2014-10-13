package gogh

import (
	"image"
	//_ "image/bmp"
	_ "image/jpeg"
	_ "image/png"
)

func (src *Mat) Grayscale() image.Image {
	img := src.src
	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	model := gray.ColorModel()
	for i := 0; i < bounds.Max.X; i++ {
		for j := 0; j < bounds.Max.Y; j++ {
			gray.Set(i, j, model.Convert(img.At(i, j)))
		}
	}
	return gray
}
func (src *Mat) Binarization(T int, reverse bool) *Mat {
	dst := src.Clone()
	bounds := src.Bounds()
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			pixel := dst.At(x, y)
			if reverse {
				if T < pixel.Gray() {
					pixel.Set(0, 0, 0)
				} else {
					pixel.Set(255, 255, 255)
				}
			} else {
				if T > pixel.Gray() {
					pixel.Set(0, 0, 0)
				} else {
					pixel.Set(255, 255, 255)
				}
			}
		}
	}
	return dst
}

func (src *Mat) Histogram() []int {
	bounds := src.Bounds()
	histogram := make([]int, 256)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := src.At(x, y).Gray() //.Gray()
			histogram[pixel]++
		}
	}
	return histogram
}