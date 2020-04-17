package imgsearch

import (
	"image"
	"math"

	_ "golang.org/x/image/bmp"
	"golang.org/x/image/draw"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

// BestMatch contains data about the search result
type BestMatch struct {
	Y          int
	X          int
	Confidence float64 // Value between 0 and 1
}

// SearchImage searches the imgB inside imgA
func SearchImage(imageA image.Image, imgB image.Image) BestMatch {
	bestMatch := BestMatch{
		Confidence: math.Inf(+1),
	}

	imgPixels := imageToPixels(imageA)
	imgHeight := len(imgPixels)
	imgWidth := len(imgPixels[0])

	subSetPixels := imageToPixels(imgB)
	subSetHeight := len(subSetPixels)
	subSetWidth := len(subSetPixels[0])

	subSetNormalized := normalizeImage(subSetPixels)
	imgNormalized := normalizeImage(imgPixels)

	subSetDense := mat.NewDense(subSetHeight, subSetWidth, subSetNormalized)
	imgDense := mat.NewDense(imgHeight, imgWidth, imgNormalized)

	for i := 0; i < imgHeight; i++ {
		if i+subSetWidth > imgWidth {
			continue
		}

		for j := 0; j < imgWidth; j++ {
			if j+subSetHeight > imgHeight {
				continue
			}

			imgSlice := imgDense.Slice(j, j+subSetHeight, i, i+subSetWidth)

			c := mat.NewDense(subSetHeight, subSetWidth, nil)
			c.Add(imgSlice, c)
			c.Sub(subSetDense, c)

			c.Apply(func(i int, j int, v float64) float64 {
				return math.Abs(v)
			}, c)

			mean := stat.Mean(c.RawMatrix().Data, nil)
			diff := mean

			if diff < bestMatch.Confidence {
				bestMatch.Confidence = diff
				bestMatch.Y = i
				bestMatch.X = j
			}
		}
	}

	// Convert the confidence score to be between 0 and 1
	bestMatch.Confidence = math.Abs((bestMatch.Confidence/16776705)*100 - 100)

	return bestMatch

}

func normalizeImage(img [][][3]float64) []float64 {
	output := make([]float64, 0)
	for _, height := range img {

		for _, pixel := range height {

			output = append(output, normalizeRGB(pixel))
		}
	}

	return output
}

func normalizeRGB(color [3]float64) float64 {
	return 0xFFFF*color[0] + 0xFF*color[1] + color[2]
}

func imageToPixels(src image.Image) [][][3]float64 {
	bounds := src.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	iaa := make([][][3]float64, height)
	srcRgba := image.NewRGBA(src.Bounds())
	draw.Copy(srcRgba, image.Point{}, src, src.Bounds(), draw.Src, nil)

	for y := 0; y < height; y++ {
		row := make([][3]float64, width)
		for x := 0; x < width; x++ {
			idxS := (y*width + x) * 4
			pix := srcRgba.Pix[idxS : idxS+4]
			row[x] = [3]float64{float64(pix[0]), float64(pix[1]), float64(pix[2])}
		}
		iaa[y] = row
	}

	return iaa
}
