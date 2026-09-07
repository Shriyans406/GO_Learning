// package main

// import "fmt"

// func main(){
// 	var intSlice = []int{1,2,3}
// 	fmt.Println(sumSlice(intSlice))

// 	var float32Slice = []float32{1,2,3}
// 	fmt.Println(sumSlice(float32Slice))
// }

// func sumSlice[T int | float32 | float64](slice []T) T{
// 	var sum T
// 	for _, v:= range slice{
// 		sum += v
// 	}
// 	return sum
// }

// // func sumFloat32Slice(slice []float32) float32{
// // 	var sum float32
// // 	for _, v:= range slice{
// // 		sum += v
// // 	}
// // 	return sum
// // }

// func isEmpty[T any](slice []T) bool{
// 	return len(slice)==0
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type contactInfo struct{
	Name string
	Email string
}

type purchaseInfo struct{
	Name string
	Price float32
	Amount int
}

func main(){
	var contacts []contactInfo=loadJSON[contactInfo]("./contacts.json")
	fmt.Printf("\n%+v", contacts)

	var purchases []purchaseInfo=loadJSON[purchaseInfo]("./purchases.json")
	fmt.Printf("\n%+v", purchases)

}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T{
	data, _ := ioutil.ReadFile(filePath)

	var loaded=[]T{}
	json.Unmarshal(data, &loaded)
	return loaded
}