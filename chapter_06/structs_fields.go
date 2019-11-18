package main

import "fmt"

type struct1 struct {
	i1 int
	f1 float32
	str string
}
func main() {
	ms := new(struct1)
	fmt.Println(*ms)
	ms.i1 = 10
	ms.f1 = 1.25
	ms.str = "张三"

	fmt.Printf("The int is %d\n",ms.i1)
	fmt.Printf("The float is %f\n",ms.f1)
	fmt.Printf("The str is %s\n",ms.str)
	fmt.Println(ms)
}
