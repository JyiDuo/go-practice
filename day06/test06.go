package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/*
type Person struct{
	Name string
	Age int
}
func(p Person)String()string{
	return fmt.Sprintf("%s的年龄%d",p.Name,p.Age)
}
func main(){
	a:=Person{"jiyiduo",19}
	fmt.Println(a)
}
*/
/*
func do(i interface{}){
	switch v:=i.(type){
	case int:
		fmt.Println("This is Int")
	case string:
		fmt.Println("String is here")
	default:
		fmt.Printf("%T is not have",v)
	}
}
func main(){
	do(1)
	do("hello")
	do(false)
}
*/
/*
func main(){
	var i interface{}="hello"
	s:=i.(string)
	fmt.Println(s)
	s,ok:=i.(string)
	fmt.Println(s,ok)
	f,ok:=i.(float64)
	fmt.Println(f,ok)
}
*/
/*
func main(){
	var i interface{}
	describe(i)
	a:=1
	i=a
	describe(i)
	var b string="hello"
	i=b
	describe(b)
}
func describe(i interface{}){
	fmt.Printf("%v,%T\n",i,i)
}*/
