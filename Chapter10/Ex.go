package main

import "fmt"

func main() {
	//c1 := make(chan string)
	//c2 := make(chan string)
	//
	//go func() {
	//	for {
	//		c1 <- "from 1"
	//		<-time.After(time.Second * 2)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		c2 <- "from 2"
	//		<-time.After(time.Second * 3)
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

	//pings := make(chan string, 1)
	//pongs := make(chan string, 1)
	//ping2(pings, "passed message")
	//pong(pings, pongs)
	//fmt.Println(<-pongs)
	c := make(chan string, 1)
	c2 := make(chan string, 1)
	c <- "hello"
	f(c)

	f2(c2)
	fmt.Println(<-c2)

}

//func ping2(pings chan<- string, msg string) {
//	pings <- msg
//}
//
//func pong(pings <-chan string, pongs chan<- string) {
//	msg := <-pings
//	pongs <- msg
//}

func f(c <-chan string) {
	fmt.Println(<-c)
}
func f2(c chan<- string) {
	c <- "xxxdq"
}
