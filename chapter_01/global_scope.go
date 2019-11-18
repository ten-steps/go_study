package main

import (
	"fmt"
	"go/types"
) // Package implementing formatted I/O.

const (
	A = iota
	A1
	A2
)
const (
	B = iota
	B1
	B2
)
func main() {
	a := 5
	b := 3
	c := cal(a,b,Add)
	fmt.Printf("函数的执行结果为 %d\n",c)
	fmt.Print(A,A1,A2)
	fmt.Print("\n")
	fmt.Print(B,B1,B2)
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	var j string
	var i = &j
	fmt.Printf("i 的数据类型为 %T",i)
}


func action(name string) {
	var _ = types.Func{}
}

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

type typeFun func(int,int) int

func cal(a,b int,fun typeFun) (result int){
	return fun(a,b)
}


