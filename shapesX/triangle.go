package shapesX

import (
	"fmt"
	"math"
)

type Triangle struct {
	Side1, Side2, Side3 float64
}

func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.Side1) * (s - t.Side2) * (s - t.Side3))
}

func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}

func NewTriangle(side1, side2, side3 float64) (*Triangle, error) {
	s := (side1 + side2 + side3) / 2
	areaSquared := s * (s - side1) * (s - side2) * (s - side3)
	if side1 <= 0 || side2 <= 0 || side3 <= 0 || areaSquared <= 0 {
		return nil, fmt.Errorf("side1 and side2 and side3 must be positive and non-zero, or the three sides with this values %.2f, %.2f, %.2f can't formed a real triangle", side1, side2, side3)
	}
	return &Triangle{side1, side2, side3}, nil
}
