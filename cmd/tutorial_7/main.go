package mian

import "fmt"

func main(){
	var p *int32=new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of p is: %v", i)
	p=&i
	*p=1
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of p is: %v", i)
	var k int32=2
	i=k


	var thing1=[5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memeory location is:%p", &thing1)
	var result [5]float64=square(thing1)
	fmt.Printf("\n the result is : %v",result)



}

func square(thing2=2 [5]float64) [5]float64{
	fmt.Printf("\n The memory location is :%p", &thing2)
	for i := range thing2{
		thing2[i]=thing2[i]*thng2[i]
	}
	return thing2
}