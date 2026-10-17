package main
/*
import(
	"fmt"
	"sync"
	"time"
)
type S struct{
	Mymutex sync.Mutex
	v map[string]int
}
func(s *S)Insert(key string){
	s.Mymutex.Lock()
	s.v[key]++
	s.Mymutex.Unlock()
}
func(s *S)Import(key string)int{
	s.Mymutex.Lock()
	defer s.Mymutex.Unlock()
	return s.v[key]
}
func main(){
	s:=S{v:make(map[string]int)}
	for i:=0;i<1000;i++{
		s.Insert("test")
	}
	time.Sleep(time.Second)
	fmt.Println(s.Import("test"))
}
*/