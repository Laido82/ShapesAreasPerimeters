package shapesX

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return math.Pow(s.side, 2)
}

func (s Square) Perimeter() float64 {
	return 4 * s.side
}

func NewSquare(side float64) (*Square, error) {
	if side <= 0 {
		return nil, fmt.Errorf("side must be positive and non-zero")
	}
	return &Square{side}, nil
}
