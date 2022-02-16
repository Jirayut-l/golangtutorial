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
	fmt.Println("!@#(ASDasd1das33333333dasdasdasd23123")
	fmt.Println("M23asdasd1dasda22222sdas231a#DA@131")
	fmt.Println("ddl1asdasasda2222sdasd12312k3")
	fmt.Println("M23#ASasdas222222d12312d10p--23")
	fmt.Println("!@#(asdasd123123")
	fmt.Println("M23#DA@1asdlpjasiopdfj31")
	fmt.Println("ddl1as;'dlpa;[sk3")
	fmt.Println("M23#as'];dlp[123--23")
	fmt.Println(h.name, h.age, h.isAdult)
	fmt.Println("Masad1231231dasdasdasd23as2")
	fmt.Println("as;'dk;'asdasdasl1kl2p;3'")
	fmt.Println("Masad1231dsadasd23123as2")
	fmt.Println("as;'dk;'lasdasdas1kl2p;3'")
}

func setAdult(h *human) {
	h.isAdult = h.age >= 18
}

func zero(xPrt *int) {
	*xPrt = 0
}
