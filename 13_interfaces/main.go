package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s interface{ area() float64 }) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
	fmt.Printf("Circle Area: %v\n", circle.area())
	fmt.Printf("Rectangle Area: %v\n", rectangle.area())

	//empty object
	obj := map[string]interface{}{}
	fmt.Println(obj)
	obj["aa"] = "string"
	fmt.Println(obj)
	obj["bb"] = 199.7735
	fmt.Println(obj)
	delete(obj, "aa")
	fmt.Println(obj)

	obj2 := map[string]interface{}{"xx": "xxx"}

	//merge 2 objs
	for k, v := range obj2 {
		obj[k] = v
	}
	fmt.Println(obj)
}
