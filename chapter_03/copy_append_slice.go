package main

import "fmt"

func main() {
	sl_form := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_form)
	fmt.Println(sl_to)
	fmt.Println(n)

	s13 := []int{4,5,6}
	s13 =  append(s13,7,8,9)
	fmt.Println(s13)
}
