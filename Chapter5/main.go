package main

import "fmt"

func main() {
	//Control Structures
	//for Loop
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//	//i += 1
	//}

	//if
	//for i := 1; i <= 10; i++ {
	//	if i%2 == 0 {
	//		fmt.Println("even")
	//	} else {
	//		fmt.Println("odd")
	//	}
	//
	//}

	//EX
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 {
	//		fmt.Println(i)
	//	}
	//}

	//EX2
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		}
	}
}
