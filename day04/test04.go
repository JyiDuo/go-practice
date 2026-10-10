package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		current := a
		a, b = b, a+b
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
func addr()func(int)int{
	sum:=0
	return func(x int)int{
		sum+=x
		return sum
	}
}
func main(){
	pos:=addr()
	for i:=0;i<6;i++{
		fmt.Println(pos(i))
	}
}
*/
/*
func compute(fn func(int,int)int)int{
	return fn(3,4)
}
func main(){
	add:=func(x,y int)int{
		return x+y
	}
	fmt.Println(add(1,2))
	fmt.Println(compute(add))
}
*/
/*
import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts:=map[string]int{}
	words:=strings.Fields(s)
	for _,word:=range words{
		counts[word]++
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
*/
/*func main(){
	var m map[string]int
	m=make(map[string]int)
	m["jiyiduo"]=1
	fmt.Println(m["jiyiduo"])
	m["jiyiduo"]=2
	fmt.Println(m["jiyiduo"])
	delete(m,"jiyiduo")
	fmt.Println(m["jiyiduo"])
	val,ok:=m["jiyiduo"]
	fmt.Printf("val:%d,ok:%v\n",val,ok)
}
*/
/*
type Vertex struct{
	x,y int
}
var m map[string]Vertex
var m1=map[string]Vertex{
"jiyiduo":Vertex{1,2},
"jiyiduo1":Vertex{3,4},
}
var m2=map[string]Vertex{
	"jiyiduo":{1,2},
	"jiyiduo1":{3,4},
}
func main(){
	m=make(map[string]Vertex)
	m["jiyiduo"]=Vertex{1,2}
	fmt.Println(m["jiyiduo"])
	fmt.Println(m1)
	fmt.Println(m2)
}*/
/*
import "golang.org/x/tour/pic"
func Pink(dx,dy int) [][]uint8{
	pink:=make([][]uint8,dy)
	for y:=0;y<dy;y++{
		pink[y]=make([]uint8,dx)
		for x:=0;x<dx;x++{
			pink[y][x]=uint8(x*y)
		}
	}
	return pink
}
func main(){
	pic.Show(Pink)
}
/*
/*
func main(){
	prims:=[]int{1,2,3,4,5,6}
	for i,v:=range prims{
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	for i:=range prims{
		prims[i]++
	}
	for _,v:=range prims{
		fmt.Printf("Value: %d\n", v)
	}
}*/

/*
func main(){
	var s []int
	printSlice(s)
	s=append(s,0)
	printSlice(s)
	s=append(s,1)
	printSlice(s)
	s=append(s,2,3,4,5)
	printSlice(s)
}
func printSlice(s[]int){
	fmt.Printf("len:%d,cap:%d,%v\n",len(s),cap(s),s)
}*/
