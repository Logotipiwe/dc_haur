package service

import (
	"dc_haur/src/pkg/gradient"
	"github.com/fogleman/gg"
	"github.com/logotipiwe/dc_go_utils/src/config"
	"image"
	"image/color"
)

func CreateImageCard(text string) (error, image.Image) {
	startColor := color.RGBA{74, 62, 255, 255}
	endColor := color.RGBA{219, 100, 255, 255}
	err, gradientImage := CreateGradient(720, 1280, startColor, endColor, 0, 0, 1, 1)
	if err != nil {
		return err, nil
	}

	err, img := putTextOnImage(gradientImage, text)
	if err != nil {
		return err, nil
	}

	return nil, img
}

func putTextOnImage(img *image.RGBA, text string) (error, image.Image) {
	dc := gg.NewContextForRGBA(img)
	h := float64(img.Rect.Max.Y)
	w := float64(img.Rect.Max.X)
	//TODO test it in dev-dc!
	if err := dc.LoadFontFace(config.GetConfig("FONTS_DIR")+"/helvetica_bold.ttf", w/25); err != nil {
		return err, nil
	}
	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, w/2, h/2, 0.5, 0.5, w*0.8, 2, gg.AlignCenter)
	dc.Stroke()
	return nil, img
}

func CreateGradient(h, w int, startColor, endColor color.RGBA, startX, startY, endX, endY float64) (error, *image.RGBA) {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	gradient.DrawLinear(img, startX, startY, endX, endY, []gradient.Stop{
		{X: 0, Col: startColor},
		{X: 1, Col: endColor},
	})
	return nil, img
}
