package service

import (
	"dc_haur/src/pkg/gradient"
	"image"
	"image/color"
)

func CreateImageCard(text string) *image.RGBA {
	startColor := color.RGBA{74, 62, 255, 255}
	endColor := color.RGBA{219, 100, 255, 255}
	return CreateGradient(720, 1280, startColor, endColor, 0, 0, 1, 1)
}

func CreateGradient(h, w int, startColor, endColor color.RGBA, startX, startY, endX, endY float64) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	gradient.DrawLinear(img, startX, startY, endX, endY, []gradient.Stop{
		{X: 0, Col: startColor},
		{X: 1, Col: endColor},
	})
	return img
}
