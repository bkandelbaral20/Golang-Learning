package main

import (
	"fmt"
	"math"
)

//define interfaces
type shape interface {
	area() float64
}

type circle struct {
	x,y, radius float64
}

type rectangle struct {
	width, height  float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) rectangle() float64 {
	return r.height * r.width
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {

	circle := circle{ x:0, y:3, radius:3 }
	//rectangle := rectangle{ height:10 , width:3 }

	fmt.Printf("Area of circle is %f\n", getArea(circle))
	//fmt.Printf("Area of rectangle is %f\n", getArea(rectangle))


}
