package main
/*
import(
	"fmt"
	"time"
	"context"
)
func main(){
	ctx,cancel:=context.WithTimeout(context.Background(),2*time.Second)
	defer cancel()
	ch:=make(chan string)
	go func(){
		time.Sleep(3*time.Second)
		select{
		case <-ctx.Done():
			fmt.Printf("程序被中断，停止")
			return
		case ch<-"请求成功":
		}
	}()
	select{
	case <-ctx.Done():
		fmt.Printf("程序超时")
	case v:=<-ch:
		fmt.Println(v)
	}
	time.Sleep(1500 * time.Millisecond)
}
*/
