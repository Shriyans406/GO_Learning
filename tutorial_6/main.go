package main

import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
	ownerInfo owner
}

type owner struct{
	name string
}

func main(){
	var myEngine gasEngine=gasEngine{mpg:25, gallons:15, ownerInfo: owner{name: "Alex"}}
	fmt.Println(myEngine.mpg)
	fmt.Println(myEngine.gallons)
	fmt.Println(myEngine.ownerInfo.name)
}
