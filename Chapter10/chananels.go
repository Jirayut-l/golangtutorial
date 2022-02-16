package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go ping(c)
	go printer(c)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err.Error())
	}

	//c1 := make(chan string)
	//c2 := make(chan string)
	//
	//go func() {
	//	for {
	//		c1 <- "from 1"
	//		time.Sleep(time.Second * 2)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		c2 <- "from 2"
	//		time.Sleep(time.Second * 3)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		select {
	//		case msg1 := <-c1:
	//			fmt.Println(msg1)
	//		case msg2 := <-c2:
	//			fmt.Println(msg2)
	//		}
	//	}
	//}()
	//
	//var input string
	//_, err := fmt.Scanln(&input)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
}

func ping(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
