package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, length float64
}

func NewRectangle(width, length float64) (*Rectangle, error) {
	if width <= 0 || length <= 0 {
		return nil, fmt.Errorf("width and length must be positive and non-zero")
	}
	return &Rectangle{width, length}, nil
}
func (r Rectangle) Area() float64 {
	return r.Width * r.length
}

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) (*Circle, error) {
	if radius <= 0 {
		return nil, fmt.Errorf("radius must be positive and non-zero")
	}
	return &Circle{radius}, nil
}
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
	Base, Height float64
}

func NewTriangle(base, height float64) (*Triangle, error) {
	if base <= 0 || height <= 0 {
		return nil, fmt.Errorf("base and height must be positive and non-zero")
	}
	return &Triangle{base, height}, nil
}
func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}

type Square struct {
	side float64
}

func NewSquare(side float64) (*Square, error) {
	if side <= 0 {
		return nil, fmt.Errorf("side must be positive and non-zero")
	}
	return &Square{side}, nil
}

func (s Square) Area() float64 {
	return math.Pow(s.side, 2)
}

type Shapes []struct {
	name  string
	shape Shape
}

func main() {
	log.Println("Server Started at localhost:8080")
	http.HandleFunc("/areas", func(w http.ResponseWriter, r *http.Request) { displayAreas(w) })
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func displayAreas(w http.ResponseWriter) {

	_, _ = fmt.Fprintln(w, "The areas of all the shapes:")
	_, _ = fmt.Fprintln(w, "============================")

	var shapes Shapes

	r, err := NewRectangle(15, 10)
	if err != nil {
		_, _ = fmt.Fprintln(w, "Can't create this rectangle : ", err)
	} else {
		shapes = append(shapes, struct {
			name  string
			shape Shape
		}{"Rectangle", r})
	}

	s, err := NewSquare(5)
	if err != nil {
		_, _ = fmt.Fprintln(w, "Can't create this square : ", err)
	} else {
		shapes = append(shapes, struct {
			name  string
			shape Shape
		}{name: "Square", shape: s})
	}

	t, err := NewTriangle(10, 10)
	if err != nil {
		_, _ = fmt.Fprintln(w, "Can't create this triangle : ", err)
	} else {
		shapes = append(shapes, struct {
			name  string
			shape Shape
		}{name: "Triangle", shape: t})
	}

	c, err := NewCircle(1)
	if err != nil {
		_, _ = fmt.Fprintln(w, "Can't create this circle : ", err)
	} else {
		shapes = append(shapes, struct {
			name  string
			shape Shape
		}{name: "Circle", shape: c})
	}
	for index, shape := range shapes {
		_, _ = fmt.Fprintf(w, "%d- The Area of %s is %.2f\n", index+1, shape.name, shape.shape.Area())
	}
}
