package main

import "fmt"

func main(){
	Min(1,2,3)
}

func Min(a ...int) {
	fmt.Printf("%T%v",a,a)
}
