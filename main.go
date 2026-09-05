package main

import "fmt"

func main() {
	var c=make(chan int)
	c <- 1
	var i= <- c
	fmt.Println(i)

	go process(c)
	for i:= range c{
		fmt.Println(i)
	}
}

func process(c chan int){
	c <- 123

	for i:=0; i<10; i++{
		c <- 1
	}
}