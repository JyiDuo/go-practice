package main

type List[T any] struct {
	next *List[T]
	data T
}

func (l *List[T]) Push(v T) {
	if l == nil {
		return
	}
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{data: v}
}

/*
import"fmt"
func Index[T comparable](s []T,x T)int{
	for i,v:=range s{
		if v==x{
			return i
		}
	}
	return -1
}
func main(){
	s:=[]int{1,2,3,4,5,6}
	fmt.Println(Index(s,1))
}*、
/*
import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func (r rot13Reader)Read(b []byte)(int,error){
	n,err:=r.r.Read(b)
	if n==0||err!=nil{
		return n,err
	}
	for i:=0;i<n;i++{
		if b[i] >= 'A' && b[i] <= 'Z' {
            // 大写字母变形
            b[i] = 'A' + (b[i]-'A'+13)%26
        } else if b[i] >= 'a' && b[i] <= 'z' {
            // 小写字母变形
            b[i] = 'a' + (b[i]-'a'+13)%26
        }
	}
	return n,err
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
*/
/*
import(
	"fmt"
	"strings"
	"io"
)
func main(){
	r:=strings.NewReader("hello,world,123456")
	b:=make([]byte,8)
	for{
		n,err:=r.Read(b)
		fmt.Printf("n=%v,err=%v,b=%v\n",n,err,b)
		fmt.Printf("b[:n]=%q\n",b[:n])
		if err==io.EOF{
			break
		}
	}
}
*/
/*
import"fmt"
import"math"
type ErrNegativeSqrt float64
func (e ErrNegativeSqrt)Error()string{
	return fmt.Sprintf("this number %v is not can be sqrt",float64(e))
}
func Sqrt(x float64)(float64,error){
	if x<0{
		return 0,ErrNegativeSqrt(x)
	}else{
		return math.Sqrt(x),nil
	}
}
func main(){
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(2))
}
*/
/*
import(
	"fmt"
	"time"
)
type Myerror struct{
	When time.Time
	What string
}
func (e*Myerror)Error()string{
	return fmt.Sprintf("at %t,%s",e.When,e.What)
}
func run()error{
	return &Myerror{
		time.Now(),
		"it do not work!",
	}
}
func main(){
	if err:=run();err!=nil{
		fmt.Println(err)
	}
}
*/
