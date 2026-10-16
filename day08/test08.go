package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	tick := time.Tick(10 * time.Millisecond)
	boom := time.After(50 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] Boom!\n", elapsed())
		default:
			fmt.Printf("[%6s]      \n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/*
import"fmt"
func feibo(c chan int,quit chan int){
	x,y:=0,1
	for{
		select{
		case c<-x:
			x,y=y,x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main(){
	quit:=make(chan int)
	c:=make(chan int)
	go func(){
		for i:=0;i<10;i++{
			fmt.Println(<-c)
		}
		quit<-0
	}()
	feibo(c,quit)
}
*/
/*
import"fmt"
func feibonaqi(n int,c chan int){
	x,y:=1,0
	for i:=0;i<n;i++{
		c<-x
		x,y=y,x+y
	}
	close(c)
}
func main(){
	c:=make(chan int,10)
	go feibonaqi(cap(c),c)
	for i:=range c{
		fmt.Println(i)
	}
}
*/
/*
import"fmt"
func main(){
	c:=make(chan int,2)
	c<-1
	c<-2
	p1:=<-c
	p2:=<-c
	fmt.Println(p1,p2)
}
*/
/*
import"fmt"
func Sum(s []int,c chan int){
	sum:=0
	for _,v:=range s{
		sum+=v
	}
	c<-sum
}
func main(){
	s:=[]int{1,1,1,1,1,1}
	c:=make(chan int)
	go Sum(s[:len(s)/2],c)
	go Sum(s[len(s)/2:],c)
	x,y:=<-c,<-c
	fmt.Println(x,y,x+y)
}
*/
/*
import(
	"fmt"
	"time"
)
func Say(s string){
	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}
func main(){
	go Say("ni hao")
	Say("wow")
}
*/
