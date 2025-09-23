package shapesX

import "fmt"

type Rectangle struct {
	Width, length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.length)
}

func NewRectangle(width, length float64) (*Rectangle, error) {
	if width <= 0 || length <= 0 {
		return nil, fmt.Errorf("width and length must be positive and non-zero")
	}
	return &Rectangle{width, length}, nil
}
