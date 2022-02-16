package main

import (
	"fmt"
	"math"
)

func main() {
	//talkers := [2]talker{
	//	human1{name: "Somchai", age: 23},
	//	parrot{name: "Dum", age: 2},
	//}
	//for _, talker := range talkers {
	//	talker.talk()
	//}

	c := Circle{x: 8, y: 6, r: 6}
	//fmt.Println(c)
	//fmt.Println(CalculationCircle(&c))
	//fmt.Println(c)
	//fmt.Println(c.area())
	r := Rectangle{l: 10, w: 20}
	fmt.Println(r.area())
	m := MultiShape{shape: []Shape{&c, &r}}
	fmt.Println("TotalArea:", totalArea(&m))
}

//#region  Type
type human1 struct {
	name string
	age  int
}
type parrot struct {
	name string
	age  int
}

type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	l, w float64
}

type MultiShape struct {
	shape []Shape
}

type talker interface {
	talk()
}

type Shape interface {
	area() float64
}

//endregion

//#region Function

func CalculationCircle(c *Circle) float64 {
	d := c.x + c.y + c.r
	c.x = 9
	return math.Pi * d
}

func (r *Rectangle) area() float64 {
	return r.l * r.w
}

func (c *Circle) area() float64 {
	d := c.x + c.y + c.r
	return math.Pi * d
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (h human1) talk() {
	fmt.Println("Human - I'm talking.")
}
func (p parrot) talk() {
	fmt.Println("Parrot  - I'm talking.")
}

func (m *MultiShape) area() float64 {
	area := 0.0
	for _, s := range m.shape {
		area += s.area()
	}
	return area
}

//#endregion
