package main

import "fmt"

func main() {
	intArr := [3]int32{1, 2, 3}
	fmt.Println(intArr)
	// intArr[1]=123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice=append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32=[]int32{8,9}
	intSlice=append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32=make([]int32, 3, 8)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice3), cap(intSlice3))

	var myMap map[string]uint8 = make (map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])

	var age, ok = myMap2["Adam"]
	delete(myMap2, "Adam")

	if ok{
		fmt.Println("The age is %v", age)
	}
	else{
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr{
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	var i int=0
	for i<10{
		if i>10{
			break
		}
		fmt.Println(i)
		i=i+1
	}

	for i:=0; i<10; i++{
		fmt.Println(i)
	}
	
}
