package main

import{
	"fmt"
	"math/rand"
	"time"

}

var wg = sync.WaitGroup{}
var dbData=[]string{"id1", "id2", "id3", "id4", "id5"}
var results=[] string{}

func main(){
	t0 := time.Now()
	for i := 0; i<len(dbData); i++{
		wg.Add(1)
		dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTime taken for sequential execution: %v", time.Since(t0))
	fmt.Printf("\nResults: %v", results)

	func dbCall(i int){
		var delay float32 = rand.Float32()*2000
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println(dbData[i])
		results=append(results, dbData[i])
		wg.Done()
	}
}