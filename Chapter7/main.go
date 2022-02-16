package main

func main() {
	//score := []float64{98, 93, 77, 82, 83}
	//fmt.Println("Result", AverageScore(score))
	//x, y := f()
	//println(x, y)

	//Closure
	//add := func(x, y int) int {
	//	return x + y
	//}
	//println(add(1, 1))
	//nextEven := makeEvenGenerator()
	//println(nextEven())
	//println(nextEven())
	//println(nextEven())
	//next := incrementor()
	//println(next())
	//println(factorial(4))

	//Defer,Panic, Recover
	//defer println("AA")
	//println("BB")
	//defer println("CC")
	//foo()
	//defer bar()
}

func foo() {
	defer println("EE")
	println("DD")
	defer println("FF")
}

func bar() {
	for i := 0; i < 3; i++ {
		println(i)
	}
}

func AverageScore(valuesList []float64) float64 {
	total := 0.0
	for _, value := range valuesList {
		total += value
	}
	return total / float64(len(valuesList))
}

func f() (int, int) {
	return 5, 6
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

//region EX
//func ex2(value int) (int, bool) {
//	result1 := 0
//	result1 = value / 2
//
//}

//#endregion
