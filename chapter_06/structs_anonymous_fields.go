package main

import "fmt"

type innerS struct {
	int1 int
	int2 int
}

type outerS struct{
	a int
	b int
	int
	innerS
}
func main() {
	outer := new(outerS)
	outer.a = 5
	outer.b =10
	outer.int = 12
	outer.int1 = 15
	outer.int2 = 18

	fmt.Printf("outer.b is: %d\n", outer.a)
	fmt.Printf("outer.c is: %f\n", outer.b)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.int1)
	fmt.Printf("outer.in2 is: %d\n", outer.int2)

	// 使用结构体字面量
	outer2 := outerS{6, 7, 60, innerS{5, 10}}
	fmt.Printf("outer2 is:%v", outer2)
}
