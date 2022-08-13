package imgproc

import (
	"image"
	"image/color"
	"image/draw"
)

func WhiteBG(src image.Image) image.Image {
	base := image.NewRGBA(src.Bounds())
	draw.Draw(base, base.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)
	draw.Draw(base, src.Bounds(), src, src.Bounds().Min, draw.Over)
	return base
}
