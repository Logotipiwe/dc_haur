package service

import (
	"dc_haur/src/pkg/gradient"
	"image"
	"image/color"
)

func CreateImageCard(text string) *image.RGBA {
	return CreateGradient(720, 1280, color.RGBA{0, 0, 0, 255},
		color.RGBA{255, 255, 255, 255}, 0, 0, 1, 1)
}

func CreateGradient(h, w int, startColor, endColor color.RGBA, startX, startY, endX, endY float64) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	gradient.DrawLinear(img, startX, startY, endX, endY, []gradient.Stop{
		{X: 0, Col: startColor},
		{X: 1, Col: endColor},
	})
	return img
}
