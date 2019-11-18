package main

import "fmt"

func main() {
	var arr [6]int
	var slice1 []int = arr[2:5]
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Printf("arr 的数据类型 %T,slice1 的数据类型为 %T\n", arr, slice1)
	fmt.Println(arr, slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1[0] = 125
	fmt.Println(slice1)
	fmt.Println(arr)
}
