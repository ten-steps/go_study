package main

import "fmt"

type TwoInts struct {
	a int
	b int
}
func main() {
	two1 := TwoInts{5,15}
	fmt.Printf("The sum is: %d\n", two1.AddThem())
}

func (this *TwoInts) AddThem() int{
	return this.a + this.b
}