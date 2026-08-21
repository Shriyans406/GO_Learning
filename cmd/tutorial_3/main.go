package main

import "fmt"

func main(){
	var printValue string = "Hello world"
	printMe(printValue)

	var numerator int = 10
	var denominator int = 2
	var result int = intDivision(numerator, denominator)
	fmt.Println(result)	
}
       
func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int)(int, int) {
	var result int= numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}