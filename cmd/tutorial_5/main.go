package main

import "fmt"

func main() {
	var myString="resume"
	var indexed=myString[0]
	//fmt.Println(indexed)
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v:=range myString{
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of the 'myString' is %v", len(myString))

	var myRune='a'
	fmt.Printf("\nmyRune=%v", myRune)
}