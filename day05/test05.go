package main

import "fmt"

type I interface {
	M()
}
type T struct {
	s string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("空")
		return
	}
	fmt.Println(t.s)
}
func main() {
	var i I
	var t *T
	i = t
	i.M()
}

/*
type Abser interface{
	abs()int
}
type Vertex struct{
	x,y int
}
func (v *Vertex)abs()int{
	return v.x*v.y
}
type Myint int
func (f Myint)abs()int{
	if f<0{
		return -int(f)
	}else{
		return int(f)
	}
}
func main(){
	var a Abser
	f:=Myint(1)
	v:=Vertex{1,2}
	a=f
	a=&v
	fmt.Println(a.abs())
}
*/
/*
type Vertex struct{
	x,y int
}
func (v Vertex)abs(f int)Vertex{
	v.x=v.x+f
	v.y=v.y+f
	return v
}
func main(){
	v:=Vertex{1,2}
	fmt.Println(v)
	fmt.Println(v.abs(1))
}
*/
/*
type Myint int
func (f Myint)Abs()int{
	if f<0{
		return -int(f)
	}else{
		return int(f)
	}
}
func main(){
	var f Myint =1
	fmt.Println(f.Abs())
}
*/
/*
import"math"
type Vertex struct{
	x,y float64
}
func (v Vertex)abs()float64{
	return math.Sqrt(v.x*v.x+v.y*v.y)
}
func main(){
	v:=Vertex{3,4}
	fmt.Println(v.abs())
}
*/
