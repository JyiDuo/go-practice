package main
import(
	"fmt"
	"time"
)
func main(){
	ch:=make(chan string)
	go func(chan string){
		time.Sleep(1*time.Second)
		ch<-"请求成功"
	}(ch)
	select{
	case v:=<-ch:
		fmt.Println(v)
	case <-time.After(2*time.Second):
		fmt.Printf("请求超时")
	}

}