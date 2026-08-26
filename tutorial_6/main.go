package main

import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
	ownerInfo owner
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

type owner struct{
	name string
}

func(e gasEngine) milesleft() uint8{
	return e.gallons*e.mpg
}

func(e electricEngine) milesleft() uint8{
	return e.kwh*e.mpkwh
}

type engine interface{
	milesleft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles<=e.milesleft(){
		fmt.Println("You can make it there!")
	}
	else{
		fmt.Println("Need to fuel up first!")
	}
}

func main(){
	var myEngine gasEngine=gasEngine{mpg:25, gallons:15, ownerInfo: owner{name: "Alex"}}
	fmt.Println(myEngine.mpg)
	fmt.Println(myEngine.gallons)
	fmt.Println(myEngine.ownerInfo.name)

	var myEngine2=struct{
		mpg uint8
		gallons uint8
	}{25, 15}

	fmt.Println(myEngine2.mpg)
	fmt.Println(myEngine2.gallons)
}
