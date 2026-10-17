package main
/*
import "golang.org/x/tour/tree"
import"fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walk(t *tree.Tree, ch chan int){
	if t==nil{
		return
	}
	walk(t.Left,ch)
	ch<-t.Value
	walk(t.Right,ch)
}
func Walk(t *tree.Tree,ch chan int){
	walk(t,ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	c1:=make(chan int)
	c2:=make(chan int)
	go Walk(t1,c1)
	go Walk(t2,c2)
	for v1:=range c1{
		v2:=<-c2
		if v1!=v2{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("开始测试 Same 函数...")

	// 测试 1：两棵一样的树
	fmt.Print("比较 tree.New(1) 和 tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("返回 true (正确)")
	} else {
		fmt.Println("返回 false (错误)")
	}

	// 测试 2：两棵不一样的树
	fmt.Print("比较 tree.New(1) 和 tree.New(2): ")
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("返回 true (错误)")
	} else {
		fmt.Println("返回 false (正确)")
	}
}*/
