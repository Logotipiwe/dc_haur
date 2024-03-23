package service

import (
	"dc_haur/src/internal/model"
	utils "dc_haur/src/pkg"
	"dc_haur/src/pkg/gradient"
	"github.com/fogleman/gg"
	config "github.com/logotipiwe/dc_go_config_lib"
	"image"
	"image/color"
	"strconv"
	"strings"
)

func CreateImageCardFromQuestion(question *model.Question, colorStart string, colorEnd string) (image.Image, error) {
	startColor := parseStringColorOrDefault(colorStart, color.RGBA{74, 62, 255, 255})
	endColor := parseStringColorOrDefault(colorEnd, color.RGBA{219, 100, 255, 255})

	imageHeight := config.GetConfigIntOr("IMAGE_HEIGHT", 720)
	imageWidth := config.GetConfigIntOr("IMAGE_WIDTH", 1280)

	gradientImage, err := createGradient(imageHeight, imageWidth, startColor, endColor, 0, 0, 1, 1)
	if err != nil {
		return nil, err
	}

	var img image.Image
	if question.AdditionalText != nil {
		img, err = putTextWithAdditionalOnImage(gradientImage, question.Text, *question.AdditionalText)
		if err != nil {
			return nil, err
		}
	} else {
		img, err = putTextOnImage(gradientImage, question.Text)
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

func putTextOnImage(img *image.RGBA, text string) (*image.RGBA, error) {
	dc := gg.NewContextForRGBA(img)
	h := float64(img.Rect.Max.Y)
	w := float64(img.Rect.Max.X)

	textY := utils.GetImageTextY()
	textWidth := utils.GetImageTextWidth()
	fontSize := utils.GetImageFontSize()

	if err := dc.LoadFontFace(config.GetConfig("FONTS_PATH"), fontSize); err != nil {
		return nil, err
	}
	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, w/2, h*textY, 0.5, 0.5, w*textWidth, 2, gg.AlignCenter)
	dc.Stroke()
	return img, nil
}

func putTextWithAdditionalOnImage(img *image.RGBA, text string, additional string) (*image.RGBA, error) {
	dc := gg.NewContextForRGBA(img)
	h := float64(img.Rect.Max.Y)
	w := float64(img.Rect.Max.X)

	additionalY := utils.GetImageAdditionalTextY()
	textY := utils.GetImageTextWithAdditionalY()
	textWidth := utils.GetImageTextWidth()
	fontSize := utils.GetImageFontSize()
	additionalFontSize := utils.GetImageAdditionalFontSize()

	if err := dc.LoadFontFace(config.GetConfig("FONTS_PATH"), fontSize); err != nil {
		return nil, err
	}
	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, w/2, h*textY, 0.5, 0.5, w*textWidth, 2, gg.AlignCenter)

	if err := dc.LoadFontFace(config.GetConfig("FONTS_PATH"), additionalFontSize); err != nil {
		return nil, err
	}
	dc.SetColor(color.White)

	dc.DrawStringWrapped(additional, w/2, h*additionalY, 0.5, 0.5, w*textWidth, 2, gg.AlignCenter)
	dc.Stroke()
	return img, nil
}

func createGradient(h, w int, startColor, endColor color.RGBA, startX, startY, endX, endY float64) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	gradient.DrawLinear(img, startX, startY, endX, endY, []gradient.Stop{
		{X: 0, Col: startColor},
		{X: 1, Col: endColor},
	})
	return img, nil
}
