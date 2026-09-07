package main

import {
	"fmt"
	"math/rand"
	"time"
}

// func main() {
// 	var c=make(chan int)
// 	c <- 1
// 	var i= <- c
// 	fmt.Println(i)

// 	go process(c)
// 	for i:= range c{
// 		fmt.Println(i)
// 	}
// }

// func process(c chan int){
// 	c <- 123

// 	for i:=0; i<10; i++{
// 		c <- 1
// 	}
// } 

var MAX_CHICKEN_PRICE=5

func main(){
	var chickenChannel=make(chan string)
	var websites=[]string{"walmart.com", "costco.com", "wholefoods.com"}
	for i:= range websites{
		go checkChickenPricea(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan string){
	for{
		time.Sleep(time.Second*1)
		var chickenPrice=rand.Float32()*20
		if chickenPrice<=MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}