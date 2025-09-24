package main

import (
	"fmt"
	"log"
	"main/methods"
	"main/shapesX"

	"github.com/gofiber/fiber/v2"
)

func main() {

	var width, height float64
	fmt.Print("Enter the values for the width and height of the rectangle : ")
	_, _ = fmt.Scanln(&width, &height)

	var side float64
	fmt.Print("Enter the value for the side of the square : ")
	_, _ = fmt.Scanln(&side)

	var side1, side2, side3 float64
	fmt.Print("Enter the values for the three sides of the triangle : ")
	_, _ = fmt.Scanln(&side1, &side2, &side3)

	var radius float64
	fmt.Print("Enter the value for the radius of the circle : ")
	_, _ = fmt.Scanln(&radius)

	rectangle, errRectangle := shapesX.NewRectangle(width, height)

	square, errSquare := shapesX.NewSquare(side)

	triangle, errTriangle := shapesX.NewTriangle(side1, side2, side3)

	circle, errCircle := shapesX.NewCircle(radius)

	log.Println("Server Started at localhost:8080")

	r := fiber.New()
	r.Get("/AreasPerimeters", func(c *fiber.Ctx) error {
		methods.DisplayAreasPerimeters(c, rectangle, square, triangle, circle, errRectangle, errSquare, errTriangle, errCircle)
		return nil
	})
	log.Fatal(r.Listen("localhost:8080"))

}
