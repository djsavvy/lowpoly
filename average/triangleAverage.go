package average

import (
	"errors"
	"image"
	"image/color"
	"math"
)

type colorAccum struct {
	sumR       uint64
	sumG       uint64
	sumB       uint64
	pixelCount uint64
}

//TriangleAverage averages the colors of all points contained in or on the triangle defined by the given vertices
func TriangleAverage(input *image.Image, output *image.RGBA, A, B, C *image.Point) (err error) {
	//degeneracy check
	slopeCheck := (A.X-B.X)*(A.Y-C.Y) - (A.X-C.X)*(A.Y-B.Y)
	if slopeCheck == 0 {
		return errors.New("input points " + A.String() + ", " + B.String() + ", " + C.String() +
			" to triangleAverage cannot be collinear")
	}

	colorAdderCreator := func() func(*image.Point) interface{} {
		sum := colorAccum{uint64(0), uint64(0), uint64(0), uint64(0)}
		return func(p *image.Point) interface{} {
			pR, pG, pB, _ := (*input).At(p.X, p.Y).RGBA()
			sum.sumR += uint64(pR)
			sum.sumG += uint64(pG)
			sum.sumB += uint64(pB)
			sum.pixelCount += uint64(1)
			return sum
		}
	}

	colorAdder := colorAdderCreator()
	colorSum := loopOverTriangle(input, A, B, C, colorAdder).(colorAccum)
	avgR := uint8((colorSum.sumR / colorSum.pixelCount) / uint64(255))
	avgG := uint8((colorSum.sumG / colorSum.pixelCount) / uint64(255))
	avgB := uint8((colorSum.sumB / colorSum.pixelCount) / uint64(255))
	avgColor := color.RGBA{avgR, avgG, avgB, 255}

	colorWriterCreator := func() func(*image.Point) interface{} {
		return func(p *image.Point) interface{} {
			output.SetRGBA(p.X, p.Y, avgColor)
			return nil
		}
	}

	colorWriter := colorWriterCreator()
	loopOverTriangle(input, A, B, C, colorWriter)

	return nil
}

//loopOverTriangle returns the output from the last point traversed (no guarantees about which point that is)
func loopOverTriangle(input *image.Image, A, B, C *image.Point, fn func(*image.Point) interface{}) (result interface{}) {
	//sort vertices by increasing Y coordinate
	if A.Y > B.Y {
		A, B = B, A
	}
	if A.Y > C.Y {
		A, C = C, A
	}
	if B.Y > C.Y {
		B, C = C, B
	}

	if A.Y == B.Y {
		height := C.Y - A.Y
		deltaX1 := float64(A.X-C.X) / float64(height)
		deltaX2 := float64(B.X-C.X) / float64(height)
		result = loopOverFlatTopTriangle(input, C, height, deltaX1, deltaX2, fn)
		return result
	}
	if B.Y == C.Y {
		height := C.Y - A.Y
		deltaX1 := float64(B.X-A.X) / float64(height)
		deltaX2 := float64(C.X-A.X) / float64(height)
		result = loopOverFlatBottomTriangle(input, A, height, deltaX1, deltaX2, fn)
		return result
	}

	//all 3 vertices have different Y coords -- have to split into flatTop and flatBottom

	height := B.Y - A.Y
	deltaX1 := float64(B.X-A.X) / float64(height)
	deltaX2 := float64(C.X-A.X) / float64(C.Y-A.Y)
	loopOverFlatBottomTriangle(input, A, height, deltaX1, deltaX2, fn)

	height = C.Y - B.Y
	deltaX1 = float64(B.X-C.X) / float64(height)
	deltaX2 = float64(A.X-C.X) / float64(C.Y-A.Y)
	//subtract 1 from height so we don't doublecount dividing line
	//don't subtract from 1 before calculating deltaX1 and deltaX2
	result = loopOverFlatTopTriangle(input, C, height-1, deltaX1, deltaX2, fn)
	return result

}

//floating point imprecision doesn't matter as long as calculations are deterministic
func loopOverFlatTopTriangle(input *image.Image, bottomVertex *image.Point, height int, deltaX1, deltaX2 float64, fn func(*image.Point) interface{}) (result interface{}) {
	if deltaX1 > deltaX2 {
		deltaX1, deltaX2 = deltaX2, deltaX1
	}
	leftBound, rightBound := float64(bottomVertex.X), float64(bottomVertex.X)
	for yVal := bottomVertex.Y; yVal >= bottomVertex.Y-height; yVal-- {
		for xVal := int(math.Ceil(leftBound)); xVal <= int(rightBound); xVal++ {
			result = fn(&image.Point{xVal, yVal})
		}
		leftBound += deltaX1
		rightBound += deltaX2
	}
	return result
}

//floating point imprecision doesn't matter as long as calculations are deterministic
func loopOverFlatBottomTriangle(input *image.Image, topVertex *image.Point, height int, deltaX1, deltaX2 float64, fn func(*image.Point) interface{}) (result interface{}) {
	if deltaX1 > deltaX2 {
		deltaX1, deltaX2 = deltaX2, deltaX1
	}
	leftBound, rightBound := float64(topVertex.X), float64(topVertex.X)
	for yVal := topVertex.Y; yVal <= topVertex.Y+height; yVal++ {
		for xVal := int(math.Ceil(leftBound)); xVal <= int(rightBound); xVal++ {
			result = fn(&image.Point{xVal, yVal})
		}
		leftBound += deltaX1
		rightBound += deltaX2
	}
	return result
}
