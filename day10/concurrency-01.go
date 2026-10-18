package main
/*
import(
	"fmt"
	"sync"
)
*/
/*
func main(){
	var wg sync.WaitGroup
	for i:=0;i<5;i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			fmt.Printf("%d个协程\n",i)
			for j:=1;j<=20;j++{
				fmt.Println(j)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("所有协程执行完毕，主程序退出")
}*/