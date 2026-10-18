package main
/*
import "fmt"
import"sync"
var wg sync.WaitGroup
func Insert(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
func Import(ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
	
}
func main(){
	ch:=make(chan int)
	go Insert(ch)
	wg.Add(1)
	go Import(ch)
	wg.Wait()
}
*/
/*
func Insert(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
func Import(ch chan int,done chan bool) {
	for v := range ch {
		fmt.Println(v)
	}
	done<-true
}
func main(){
	ch:=make(chan int)
	bone:=make(chan bool)
	go Insert(ch)
	go Import(ch,bone)
	<-bone
}
*/