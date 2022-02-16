package main

import "fmt"

func main() {

	//var x string
	//x = "Hello World By Variable"
	//fmt.Println(x)
	//testText := 5
	//fmt.Println(testText)

	//TestString
	//dogsName := "Max"
	//fmt.Println("My dog's Name is", dogsName)
	//
	////Defining Multiple Variables
	//fmt.Print("Enter a number:")
	//var input float64
	//f, err := fmt.Scanf("%f", &input)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//f = int(input)
	//output := f * 2
	//fmt.Println(output)
	//

	//EX: Convert Fahrenheit : Celsius
	fmt.Print("Enter a Fahrenheit:")
	var inputValue int
	fahrenheit, err := fmt.Scanf("%d", &inputValue)
	if err != nil {
		fmt.Println(fahrenheit, err)
	}
	celsius := float64((inputValue - 32) * 5 / 9)
	fmt.Println("Fahrenheit Value", inputValue)
	fmt.Println("Result", celsius)
}
