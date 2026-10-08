package main

import "fmt"

func main() {
	x := 10
	defer fmt.Println("修改前：", x)
	x = 20
	fmt.Println("修改后：", x)
}

/*
import "runtime"
import "time"
func main(){
	switch os:=runtime.GOOS;os{
	case "darwin":
		fmt.Println("MacOs")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Unknown OS")
	}

	fmt.Println("When's Saturday?")
	today:=time.Now().Weekday()
	switch time.Saturday{
		case today+0:
			fmt.Println("Today is Saturday")
		case today+1:
			fmt.Println("Tomorrow")
		case today+2:
			fmt.Println("In two days")
		default:
			fmt.Println("After two days")
	}

}
*/
/*
func Sqrt(x float64) float64 {
	z:=1.0
	for i:=0;i<10;i++{
		z-=(z*z-x)/(2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}*/

/*
func add(x,y int)int{
	return x+y
}
func main(){
	if v:=add(1,2);v>0{
		fmt.Print(v)
	}else {
		fmt.Print("v is not greater than 0")
	}
}*/

/*
func main(){
	sum:=0
	for i:=0;i<5;i++{
		sum+=i
	}
	j:=0
	for j<5{
		sum+=j
		j++
	}
	fmt.Print(sum)
}
*/
/*
const Pi=3.14
func main(){
	fmt.Println(Pi)
	const world="World"
	fmt.Println("Hello",world)
}
*/
/*
func main(){
	i:=1
	f:=float64(i)
	k:=int(f)
	fmt.Printf("Type:%T,Value:%v\n",i,i)
	fmt.Printf("Type:%T,Value:%v\n",f,f)
	fmt.Printf("Type:%T,Value:%v\n",k,k)
}*/

/*
var(
	c bool =true
	j int =1
	i float64 =3.14
)
func main(){
	fmt.Printf("Type:%T,Value:%v\n",c,c)
	fmt.Printf("Type:%T,Value:%v\n",j,j)
	fmt.Printf("Type:%T,Value:%v\n",i,i)
}
*/

/*
func main(){
	var i ,j int=1,2
	k:=3
	c,python,java:=true,true,"no"
	fmt.Println(i,j,k,c,python,java)
}
*/
