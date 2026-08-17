package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var intNum int=2323
	fmt.Println(intNum)

	var floatNum float64=12345.8
	fmt.Println(floatNum)

	var floatNum32 float32=10.2
	var intNum32 int32=2
	var result float32=floatNum32+float32(intNum32)
	fmt.Println(result)

	var intNum1 int=3
	var intNum2 int=2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string="Hello"+" "+"World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("Y"))

	var myRune rune='a'
	fmt.Println(myRune)

	var myBoolean bool=false
	fmt.Println(myBoolean)

	// myVar :- "text"
	// fmt.Println(myVar)

}