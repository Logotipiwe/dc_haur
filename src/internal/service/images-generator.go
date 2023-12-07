package service

import (
	"image"
	"image/color"
	"image/png"
	"net/http"
)

type ImageGenerator struct {
}

func (g *ImageGenerator) HandleGradientRequest(w http.ResponseWriter, r *http.Request) {
	width, height := 1280, 720
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	startColor := color.RGBA{74, 62, 255, 255}
	endColor := color.RGBA{219, 100, 255, 255}
	drawGradient(img, startColor, endColor)
	w.Header().Set("Content-Type", "image/png")
	err := png.Encode(w, img)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func drawGradient(img *image.RGBA, startColor, endColor color.RGBA) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		// Calculate the interpolation factor based on the y-coordinate
		t := float64(y-bounds.Min.Y) / float64(bounds.Max.Y-bounds.Min.Y)

		// Interpolate between startColor and endColor
		r := uint8(float64(startColor.R)*(1-t) + float64(endColor.R)*t)
		g := uint8(float64(startColor.G)*(1-t) + float64(endColor.G)*t)
		b := uint8(float64(startColor.B)*(1-t) + float64(endColor.B)*t)
		a := uint8(float64(startColor.A)*(1-t) + float64(endColor.A)*t)

		// Set the color for each pixel in the current row
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			img.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}
}

/*

func handleGradientRequest(w http.ResponseWriter, r *http.Request) {
	width, height := 800, 600

	// Get start and end points from query parameters (default to top-left and bottom-right)
	startX, startY := getFloatQueryParam(r, "startX", 0.0), getFloatQueryParam(r, "startY", 0.0)
	endX, endY := getFloatQueryParam(r, "endX", 1.0), getFloatQueryParam(r, "endY", 1.0)

	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Define gradient colors
	startColor := color.RGBA{255, 0, 0, 255} // Red
	endColor := color.RGBA{0, 0, 255, 255}   // Blue

	// Draw gradient background
	drawGradient(img, startColor, endColor, startX, startY, endX, endY)

	// Set the Content-Type header
	w.Header().Set("Content-Type", "image/png")

	// Encode the image and write it to the response writer
	err := png.Encode(w, img)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func drawGradient(img *image.RGBA, startColor, endColor color.RGBA, startX, startY, endX, endY float64) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Calculate the interpolation factors based on the x and y coordinates
			tX := float64(x-bounds.Min.X) / float64(bounds.Max.X-bounds.Min.X)
			tY := float64(y-bounds.Min.Y) / float64(bounds.Max.Y-bounds.Min.Y)

			// Interpolate between startColor and endColor
			r := uint8(float64(startColor.R)*(1-tX)*(1-tY) + float64(endColor.R)*tX*tY)
			g := uint8(float64(startColor.G)*(1-tX)*(1-tY) + float64(endColor.G)*tX*tY)
			b := uint8(float64(startColor.B)*(1-tX)*(1-tY) + float64(endColor.B)*tX*tY)
			a := uint8(float64(startColor.A)*(1-tX)*(1-tY) + float64(endColor.A)*tX*tY)

			img.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}
}

func getFloatQueryParam(r *http.Request, key string, defaultValue float64) float64 {
	value := r.URL.Query().Get(key)
	if value == "" {
		return defaultValue
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("Invalid float value for %s. Using default.", key)
		return defaultValue
	}
	return floatValue
}
*/
