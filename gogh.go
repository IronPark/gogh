// gogh project gogh.go
package gogh

import (
	//"fmt"
	"github.com/suapapa/go_hangul"
	"image"
	"image/color"
)

<<<<<<< HEAD
func NewMat(src image.Image) *Mat {
	dst := ImageToRGBA(src)
	return &Mat{dst}
=======
var (
	GrayArray = make([]color.Color, 256)
)

func NewImg(rect image.Rectangle) *Img {
	return &Img{image.NewRGBA(rect)}
>>>>>>> 445fcbcce5485b480a6c6f8fa72994eadb0a16f2
}

type Img struct {
	src *image.RGBA
}

type Pixel struct {
	src   *image.RGBA
	X     int
	Y     int
	color color.Color
}

func (src *Img) At(x, y int) *Pixel {
	return &Pixel{src.src, x, y, src.src.At(x, y)}
}

func (src *Pixel) RGBA() (int, int, int, int) {
	r, g, b, a := src.color.RGBA()
	return int(r >> 8), int(g >> 8), int(b >> 8), int(a >> 8)
}

func (src *Img) Save(path string) {
	Save(path, src.src)
}

func (src *Img) Clone() *Img {
	return clone(src)
}

func (src *Pixel) Gray() int {
	gray, _, _, _ := src.color.RGBA()
	return int(gray >> 8)
}

func (src *Pixel) Set(r, g, b int) {
	src.src.Set(src.X, src.Y, color.Color(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}))
}

func (src *Img) Bounds() image.Rectangle {
	return src.src.Bounds()
}

func (src *Img) Pixels() []uint8 {
	return src.src.Pix
}
