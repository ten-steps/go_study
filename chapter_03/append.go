package main

import "fmt"

func main() {
	var a = []int{1,2,3,4,5,6}
	b:=a[0:3]
	c:=make([]string,2,3)
	fmt.Println(b)
	fmt.Println(len(c),cap(c),c)
	d := append(b,9,10,23,25)
	println(d)
	println(b)
	println(&a)
	fmt.Println(len(d),cap(d),d)
}
