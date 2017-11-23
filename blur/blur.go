package blur

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"math"
	"sync"
)

/* TODO: implement both exact Gaussian blur and approximation with successive box blurs
 */

/*GaussianBlur blurs the image. If isExact is true, it uses an exact Gaussian kernel;
otherwise it approximates with successive box blurs.
(see http://web.csse.uwa.edu.au/research/?a=826172 and http://blog.ivank.net/fastest-gaussian-blur.html)
It automatically guesses the necessary kernel size from the value of sigma, which must be
a positive real number. */
func GaussianBlur(input *image.Image, sigma float64, isExact bool) (*image.RGBA, error) {
	if sigma <= 0 {
		return nil, errors.New("sigma must be a positive real number")
	}

	if isExact {
		return exactGaussianBlur(input, sigma), nil
	}
	return approxGaussianBlur(input, sigma), nil
}

func exactGaussianBlur(input *image.Image, sigma float64) *image.RGBA {
	kernelRadius := int(math.Ceil(3 * sigma))
	kernel := *calculateOneDimGaussianKernel(kernelRadius, sigma)
	fmt.Println(kernel)
	bounds := (*input).Bounds()
	xMin := bounds.Min.X
	xMax := bounds.Max.X
	yMin := bounds.Min.Y
	yMax := bounds.Max.Y

	var rowsWG sync.WaitGroup
	processedRowsOnly := image.NewRGBA(bounds)
	for y := yMin; y < yMax; y++ {
		go func(y int) {
			rowsWG.Add(1)
			defer rowsWG.Done()

			convR := make([]float64, xMax-xMin)
			convG := make([]float64, xMax-xMin)
			convB := make([]float64, xMax-xMin)

			for x := xMin; x < xMax; x++ {
				for i := 0; i <= kernelRadius; i++ {
					inputPixelR, inputPixelG, inputPixelB, _ := (*input).At(x+i, y).RGBA()
					convR[x-xMin] += float64(inputPixelR/255) * kernel[i]
					convG[x-xMin] += float64(inputPixelG/255) * kernel[i]
					convB[x-xMin] += float64(inputPixelB/255) * kernel[i]
				}
				for i := 1; i <= kernelRadius; i++ {
					inputPixelR, inputPixelG, inputPixelB, _ := (*input).At(x-i, y).RGBA()
					convR[x-xMin] += float64(inputPixelR/255) * kernel[i]
					convG[x-xMin] += float64(inputPixelG/255) * kernel[i]
					convB[x-xMin] += float64(inputPixelB/255) * kernel[i]
				}
			}

			borderScaleFactor := float64(1)
			for i := 0; i < kernelRadius; i++ {
				borderScaleFactor -= kernel[kernelRadius-i]
				convR[xMin+kernelRadius-i] /= borderScaleFactor
				convG[xMin+kernelRadius-i] /= borderScaleFactor
				convB[xMin+kernelRadius-i] /= borderScaleFactor
				convR[xMax-kernelRadius-1+i] /= borderScaleFactor
				convG[xMax-kernelRadius-1+i] /= borderScaleFactor
				convB[xMax-kernelRadius-1+i] /= borderScaleFactor
			}

			for i := range convR {
				processedRowsOnly.SetRGBA(xMin+i, y,
					color.RGBA{
						R: uint8(convR[i]),
						G: uint8(convG[i]),
						B: uint8(convB[i]),
						A: uint8(255),
					})
			}
		}(y)
	}
	rowsWG.Wait()

	var colsWG sync.WaitGroup
	processedRowsCols := image.NewRGBA(bounds)
	for x := xMin; x < xMax; x++ {
		go func(x int) {
			colsWG.Add(1)
			defer colsWG.Done()

			convR := make([]float64, yMax-yMin)
			convG := make([]float64, yMax-yMin)
			convB := make([]float64, yMax-yMin)

			for y := yMin; y < yMax; y++ {
				for i := 0; i <= kernelRadius; i++ {
					inputPixelR, inputPixelG, inputPixelB, _ := (*processedRowsOnly).At(x, y+i).RGBA()
					convR[y-yMin] += float64(inputPixelR/255) * kernel[i]
					convG[y-yMin] += float64(inputPixelG/255) * kernel[i]
					convB[y-yMin] += float64(inputPixelB/255) * kernel[i]
				}
				for i := 1; i <= kernelRadius; i++ {
					inputPixelR, inputPixelG, inputPixelB, _ := (*processedRowsOnly).At(x, y-i).RGBA()
					convR[y-yMin] += float64(inputPixelR/255) * kernel[i]
					convG[y-yMin] += float64(inputPixelG/255) * kernel[i]
					convB[y-yMin] += float64(inputPixelB/255) * kernel[i]
				}
			}

			borderScaleFactor := float64(1)
			for i := 0; i < kernelRadius; i++ {
				borderScaleFactor -= kernel[kernelRadius-i]
				convR[yMin+kernelRadius-i] /= borderScaleFactor
				convG[yMin+kernelRadius-i] /= borderScaleFactor
				convB[yMin+kernelRadius-i] /= borderScaleFactor
				convR[yMax-kernelRadius-1+i] /= borderScaleFactor
				convG[yMax-kernelRadius-1+i] /= borderScaleFactor
				convB[yMax-kernelRadius-1+i] /= borderScaleFactor
			}

			for i := range convR {
				processedRowsCols.SetRGBA(x, yMin+i,
					color.RGBA{
						R: uint8(convR[i]),
						G: uint8(convR[i]),
						B: uint8(convB[i]),
						A: uint8(255),
					})
			}
		}(x)
	}
	colsWG.Wait()

	return processedRowsCols
}

// outputs sum to 1
func calculateOneDimGaussianKernel(kernelRadius int, sigma float64) *[]float64 {
	result := make([]float64, kernelRadius+1)
	denominator := sigma * math.Sqrt2 * math.Sqrt(math.Pi)
	expDenom := float64(-2) * sigma * sigma
	sum := float64(0)
	for i := range result {
		result[i] = math.Exp(float64(i)*float64(i)/expDenom) / denominator
		sum += float64(2) * result[i]
	}
	sum -= result[0]
	for i := range result {
		result[i] /= sum
	}
	return &result
}

func approxGaussianBlur(input *image.Image, sigma float64) *image.RGBA {
	return nil
}
