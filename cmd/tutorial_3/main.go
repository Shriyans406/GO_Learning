package main

import (
	"errors"
	"fmt"
)

func main(){
	var printValue string = "Hello world"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 2
	var result int
	var remainder int
	result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(result)
	fmt.Println(remainder)
	fmt.Printf("The result of %d divided by %d is %d with a remainder of %d\n", numerator, denominator, result, remainder)
}
       
func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int)(int, int) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}
	
	var result int= numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}