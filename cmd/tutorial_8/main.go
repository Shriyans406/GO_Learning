package main

import{
	"fmt"
	"math/rand"
	"time"
	"sync"
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

	// func dbCall(i int){
	// 	var delay float32 = rand.Float32()*2000
	// 	time.Sleep(time.Duration(delay) * time.Millisecond)
	// 	fmt.Println(dbData[i])
	// 	results=append(results, dbData[i])
	// 	wg.Done()
	// }

	func dbCall(i int){
		var delay float32= 2000
		time.Sleep(time.Duration(delay)*time.Millisecond)
		fmt.Println("The result from the database is :", dbData[i])
		m.Lock()
		results=append(results, dbData[i])
		m.Unlock()
		wg.Done()
	}

	func save(result string){
		m.Lock()
		results=append(results, result)
		m.Unlock()
	}

	func log(){
		m.RLock()
		fmt.Println("The results are :", results)
		m.RUnlock()
	}
}