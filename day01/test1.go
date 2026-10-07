package main

/*
import (
	"fmt"
	"math/rand"
)


func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
*/

import"fmt"
func add(x int,y int)int{
	return x+y
}
func add2(x,y int)int{
	return x+y
}
func swap(x,y int)(int ,int){
	return y,x
}
func split(sum int)(x,y int){
	x=sum*4/9
	y=sum-x
	return
}
var c,python,java bool
var j int =1
func main(){
	fmt.Println(add(11,22))
	fmt.Println(add2(33,44))
	fmt.Println(swap(2026,2027))
	fmt.Println(split(17))
	var i int
	var boy,girl=true,false
	fmt.Println(j,boy,girl)
	fmt.Println(i,c,python,java)
}