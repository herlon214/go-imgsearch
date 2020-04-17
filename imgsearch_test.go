package imgsearch

import (
	"image"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeRGB(t *testing.T) {
	white := [3]float64{255, 255, 255}
	assert.Equal(t, 1.6776705e+07, normalizeRGB(white))

	black := [3]float64{0, 0, 0}
	assert.Equal(t, 0.0, normalizeRGB(black))
}

func TestSearchImage(t *testing.T) {
	imgA := loadImage("testdata/resized-poe_stash.bmp")
	imgB := loadImage("testdata/resized-chaos2.bmp")

	result := SearchImage(imgA, imgB)
	assert.Equal(t, 263, result.Y)
	assert.Equal(t, 133, result.X)
	assert.Equal(t, 95.54177759180563, result.Confidence)
}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}

	return img
}
