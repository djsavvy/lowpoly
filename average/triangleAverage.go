package average

import (
	"errors"
	"image"
	//	"image/color"
	"sort"
)

func triangleAverage(input *image.Image, A *image.Point, B *image.Point, C *image.Point, output *image.Image) (err error) {
	//collinearity check
	slopeCheck := (A.X-B.X)*(A.Y-C.Y) - (A.X-C.X)*(A.Y-B.Y)
	if slopeCheck == 0 {
		return errors.New("input points " + A.String() + ", " + B.String() + ", " + C.String() + " to triangleAverage cannot be collinear")
	}

	return nil

}

func loopOverTriangle(input *image.Image, A *image.Point, B *image.Point, C *image.Point, fn func(*image.Point) (uint64, uint64, uint64)) {

}
