package blur

import (
	"errors"
	"image"
	"math"
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
		kernelRadius := int(math.Ceil(3 * sigma))
		return exactGaussianBlur(input, kernelRadius, sigma), nil
	}
	return approxGaussianBlur(input, sigma), nil
}

func exactGaussianBlur(input *image.Image, kernelRadius int, sigma float64) *image.RGBA {
	return nil
}

func approxGaussianBlur(input *image.Image, sigma float64) *image.RGBA {
	return nil
}
