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

	//Zero Value
	var xx1 int  //default value
	var xx2 bool //default value
	x1 := 0
	x2 := false
	fmt.Println(x1, x2, xx1, xx2)

	//if statement
	point := 50
	if point >= 50 && point <= 100 {
		fmt.Println("more than 50")
	} else if point <= 20 {
		fmt.Println("less than 20")
	}

	// array
	x323 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(x323)
	//slice array  = list
	x324 := []int{1, 2, 3}
	x324 = append(x324, 4) //==add

	//anonymous function
	xAntony := func(a, b int) int {
		return a + b
	}
	sumXa := xAntony(10, 20)
	println(sumXa)
	call(add)
	call(xAntony)
	//slice sent to function
	v1 := []int{1, 2, 3}
	s1 := sum(v1)
	_ = s1
	v2 := sum1(1, 2, 3, 4, 5)
	_ = v2

}

func call(f func(int, int) int) {
	sum := f(50, 10)
	println(sum)
}

func add(a, b int) int {
	return a + b
}

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func sum1(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
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
<<<<<<< HEAD
=======
	fmt.Println("!@#(asdasd123123")
	fmt.Println("M23#DA@1asdlpjasiopdfj31")
	fmt.Println("ddl1as;'dlpa;[sk3")
	fmt.Println("M23#as'];dlp[123--23")
	fmt.Println(h.name, h.age, h.isAdult)
	fmt.Println("as;'dk;'l1kl2p;3'")
>>>>>>> a570837 (test mergeX13)
}

func setAdult(h *human) {
	h.isAdult = h.age >= 18
}

func zero(xPrt *int) {
	*xPrt = 0
}
