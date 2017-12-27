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

// to sort the points of the triangle
type byYthenX []*image.Point

func (a byYthenX) Len() int {
	return len(a)
}

func (a byYthenX) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byYthenX) Less(i, j int) bool {
	if a[i].Y < a[j].Y {
		return true
	} else if a[i].Y > a[j].Y {
		return false
	} else {
		return a[i].X < a[j].X
	}
}

func loopOverTriangle(input *image.Image, A *image.Point, B *image.Point, C *image.Point, fn func(*image.Point) (uint64, uint64, uint64)) {
	//first order the points
	vertices := []*image.Point{A, B, C}
	sort.Sort(byYthenX(vertices))

}
