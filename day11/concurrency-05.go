package main
import(
	"fmt"
	"sync"
	"time"
)
func Worker(id int,jobs<-chan int,result chan<-int,wg*sync.WaitGroup){
	defer wg.Done()
	for job :=range jobs{
		fmt.Printf("工人：%d抢到任务%d\n",id,job)
		time.Sleep(1*time.Second)
		result<-job*job
	}
}
func main(){
	jobs:=make(chan int,100)
	result:=make(chan int,100)
	var wg sync.WaitGroup
	for w:=1;w<=5;w++{
		wg.Add(1)
		go Worker(w,jobs,result,&wg)
	}
	for j := 1; j <= 50; j++ {
		jobs <- j
	}
	close(jobs)
	go func(){
		wg.Wait()
		close(result)
	}()
	for rs:=range result{
		fmt.Println(rs)
	}
	fmt.Println("所有任务处理完毕！")
}