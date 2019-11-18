package main

import "fmt"

// 命名冲突
type TypeA struct {
	a int
}
type TypeB struct {
	TypeA
	a string
}

type TypeC struct {
	TypeA
	TypeB
	a int
}
func main() {
	tt := new(TypeC)
	tt.a = 5
	tt.TypeB.a = "test"
	tt.TypeA.a = 12
	fmt.Printf("%v",tt)
}
