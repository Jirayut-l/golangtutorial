package main

import "fmt"

func main() {
	//var x [5]float64 x[0] = 98 x[1] = 93 x[2] = 77 x[3] = 82 x[4] = 83
	//x := [5]float64{98, 93, 77, 82, 83}
	//var total float64 = 0
	//for i := 0; i < len(x); i++ {
	//	total += x[i]
	//}
	//total = total / float64(len(x))
	//fmt.Println(total)

	//x2 := x[0:2]
	//fmt.Println(x2)

	//slice := []int{1, 2, 3}
	//slice2 := append(slice, 4, 5)
	//fmt.Println(slice, slice2)

	//slice1 := []int{1, 2, 3}
	//slice2 := make([]int, 2)
	//copy(slice2, slice1)
	//fmt.Println(slice1, slice2)

	//Map
	//elements := make(map[string]string) elements["H"] = "Hydrogen"
	//elements["He"] = "Helium" elements["Li"] = "Lithium" elements["Be"] =
	//"Beryllium" elements["B"] = "Boron" elements["C"] = "Carbon" elements["N"] =
	//"Nitrogen" elements["O"] = "Oxygen" elements["F"] = "Fluorine" elements["Ne"]= "Neon"
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	//delete(elements, "C")
	//if name, success := elements["H"]; true {
	//	fmt.Println(name, success)
	//}
	name, ok := elements["H"]
	if ok {
		fmt.Println(name, ok)
	}

	var elementsX2 = map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
	}

	if name, status := elementsX2["He"]["name"]; true {
		fmt.Println(name, status)
	}

	//EX
	//xEx := [6]string{"a", "b", "c", "d", "e", "f"}
	//fmt.Println(xEx[4])
	//x3 := xEx[2:5]
	//fmt.Println(x3)

	//x34 := []int{
	//	48, 96, 86, 68,
	//	57, 82, 63, 70,
	//	37, 34, 83, 27,
	//	19, 97, 9, 17,
	//}

	//x5 := make([]int, 0, 9)
	//for i := 0; i < 10; i++ {
	//	x5 = append(x5, i)
	//	fmt.Printf("cap %v, len %v, %p\n", cap(x5), len(x5), x5)
	//}

	//min := x34[0]
	//for i := 0; i < len(x34); i++ {
	//	if min > x34[i] {
	//		min = x34[i]
	//	}
	//}
	//fmt.Println(min)
}
