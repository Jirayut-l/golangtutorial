package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println("Hello, my name is TuMz")

	//var x = 0
	//x2 := "Sdasd"
	//println(x, x2)
	//
	////Array
	//names := []string{"x1", "x2", "x3"}
	//fmt.Println(names)
	//names = append(names, "x4")
	//
	////for array range
	//for _, name := range names {
	//	fmt.Println(name)
	//}
	// Function
	//fullName := printFullName("Jirayut", "Laorpongphruek")
	//fmt.Println(fullName)

	//result, err := divide(5, 0)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(result)

	//Structs
	// somChai := human{name: "Jirayut", age: 27}
	//somChai.printInfo()
	//setAdult(&somChai)
	//fmt.Println(somChai)

	x := 8
	zero(&x)
	fmt.Println(x)
}

func printFullName(firstName string, lastName string) string {
	return firstName + lastName
}

func divide(dividend float32, divisor float32) (float32, error) {
	if divisor == 0.0 {
		err := errors.New("Division by zero!")
		return 0.0, err
	}
	return dividend / divisor, nil
}

type human struct {
	name    string
	age     int
	isAdult bool
}

func (h human) printInfo() {
	fmt.Println(h.name, h.age, h.isAdult)
	fmt.Println("TestMerg2")
	fmt.Println(h.name, h.age, h.isAdult)
}

func setAdult(h *human) {
	h.isAdult = h.age >= 18
}

func zero(xPrt *int) {
	*xPrt = 0
}
