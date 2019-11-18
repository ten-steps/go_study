package main

import "fmt"

func main(){
	var il = 5
	fmt.Printf("il 的内存指针地址为 %d,%p\n",il,&il)
	var intP *int
	intP = &il
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}