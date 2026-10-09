package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}

/*
func main(){
	arr:=[6]int{1,2,3,4,5,6}
	fmt.Println(arr)
	s1:=arr[1:4]
	fmt.Printf("s1: len=%d cap=%d %v\n", len(s1), cap(s1), s1)
	s2:=s1[0:2]
	fmt.Printf("s2: len=%d cap=%d %v\n", len(s2), cap(s2), s2)
	s3:=make([]int,0,5)
	fmt.Printf("s3: len=%d cap=%d %v\n", len(s3), cap(s3), s3)
}
*/

/*
func main(){
	s:=[]int{1,2,3,4,5,6}
	fmt.Print("len:%d,cap:%d,%v\n",len(s),cap(s),s)
	s=s[:4]
	fmt.Print("len:%d,cap:%d,%v\n",len(s),cap(s),s)
	s=s[2:]
	fmt.Print("len:%d,cap:%d,%v\n",len(s),cap(s),s)
}*/

/*
func main(){
	s:=[]int{1,2,3,4,5,6}
	fmt.Println(s)
	s=s[1:4]
	fmt.Println(s)
	s=s[:2]
	fmt.Println(s)
	s=s[1:]
	fmt.Println(s)
}*/

/*
func main(){
	prims:=[]int{2,3,5,7,11,13}
	fmt.Println(prims)
	s:=[]struct{
		i int
		b bool
	}	{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
*/
/*
func main(){
	prims:=[6]int{1,2,3,4,5,6}
	var s []int=prims[1:4]
	fmt.Println(s)
	s[0]=100
	s[1]=200
	s[2]=300
	fmt.Println(s)
	fmt.Println(prims)
}*/

/*
func main(){
	var a[2]string
	a[0]="Hello"
	a[1]="World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	prims:=[6]int{1,2,3,4,5,6}
	fmt.Println(prims)
}*/

/*
type Vertex struct{
	X int
	Y int
}
var(
	v1=Vertex{1,2}
	v2=Vertex{X:1}
	v3=Vertex{}
	p=&Vertex{Y:1}
)
func main(){
	fmt.Println(v1,v2,v3,p)
}*/
/*
func main(){
	fmt.Println(Vertex{1,2})
	v:=Vertex{2,3}
	fmt.Println(v.X,v.Y)
	v.X=4
	fmt.Println(v.X,v.Y)
	p:=&v
	p.Y=4
	fmt.Println(v.X,v.Y)
}*/
/*
func main(){
	i:=1
	p:=&i
	fmt.Println(i)
	fmt.Println(*p)
	*p=2
	fmt.Println(i)
	fmt.Println(*p)
}*/
