package service

import (
	"dc_haur/src/pkg/gradient"
	"github.com/fogleman/gg"
	config "github.com/logotipiwe/dc_go_config_lib"
	"image"
	"image/color"
	"strconv"
	"strings"
)

func CreateImageCard(text string, colorStart string, colorEnd string) (image.Image, error) {
	startColor := parseStringColorOrDefault(colorStart, color.RGBA{74, 62, 255, 255})
	endColor := parseStringColorOrDefault(colorEnd, color.RGBA{219, 100, 255, 255})

	gradientImage, err := CreateGradient(720, 1280, startColor, endColor, 0, 0, 1, 1)
	if err != nil {
		return nil, err
	}

	img, err := putTextOnImage(gradientImage, text)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func parseStringColorOrDefault(colorStr string, defaultValue color.RGBA) color.RGBA {
	split := strings.Split(colorStr, ",")
	if len(split) != 3 {
		return defaultValue
	}
	red, err := strconv.Atoi(split[0])
	green, err := strconv.Atoi(split[1])
	blue, err := strconv.Atoi(split[2])
	if err != nil {
		return defaultValue
	}
	return color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
}

func putTextOnImage(img *image.RGBA, text string) (image.Image, error) {
	dc := gg.NewContextForRGBA(img)
	h := float64(img.Rect.Max.Y)
	w := float64(img.Rect.Max.X)
	// TODO test it in dev-dc!
	if err := dc.LoadFontFace(config.GetConfig("FONTS_PATH"), w/25); err != nil {
		return nil, err
	}
	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, w/2, h/2, 0.5, 0.5, w*0.8, 2, gg.AlignCenter)
	dc.Stroke()
	return img, nil
}

func CreateGradient(h, w int, startColor, endColor color.RGBA, startX, startY, endX, endY float64) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	gradient.DrawLinear(img, startX, startY, endX, endY, []gradient.Stop{
		{X: 0, Col: startColor},
		{X: 1, Col: endColor},
	})
	return img, nil
}
