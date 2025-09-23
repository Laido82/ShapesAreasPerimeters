package methods

import (
	"fmt"
	"main/shape"

	"github.com/gofiber/fiber/v2"
)

func DisplayAreasPerimeters(c *fiber.Ctx, rectangle, square, triangle, circle shape.Shape, errRectangle, errSquare, errTriangle, errCircle error) {

	_, _ = fmt.Fprintln(c, "The areas and perimeters of all the shapes:")
	_, _ = fmt.Fprintln(c, "===========================================")
	_, _ = fmt.Fprintln(c)

	type Shapes []struct {
		name  string
		shape shape.Shape
	}

	var shapes Shapes

	if errRectangle != nil {
		_, _ = fmt.Fprintln(c, "Can't create this rectangle : ", errRectangle)
	} else {
		shapes = append(shapes, struct {
			name  string
			shape shape.Shape
		}{"Rectangle", rectangle})
	}

	if errSquare != nil {
		_, _ = fmt.Fprintln(c, "Can't create this square : ", errSquare)
	} else {
		shapes = append(shapes, struct {
			name  string
			shape shape.Shape
		}{name: "Square", shape: square})
	}

	if errTriangle != nil {
		_, _ = fmt.Fprintln(c, "Can't create this triangle : ", errTriangle)
	} else {
		shapes = append(shapes, struct {
			name  string
			shape shape.Shape
		}{name: "Triangle", shape: triangle})
	}

	if errCircle != nil {
		_, _ = fmt.Fprintln(c, "Can't create this circle : ", errCircle)
	} else {
		shapes = append(shapes, struct {
			name  string
			shape shape.Shape
		}{name: "Circle", shape: circle})
	}

	for index, thisShape := range shapes {
		_, _ = fmt.Fprintf(c, "%d- The Area of %s is %.2f\n", index+1, thisShape.name, thisShape.shape.Area())
		_, _ = fmt.Fprintf(c, "  The Perimeter of %s is %.2f\n", thisShape.name, thisShape.shape.Perimeter())
		_, _ = fmt.Fprintln(c)
	}
}
