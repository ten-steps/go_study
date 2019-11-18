package main

import "fmt"

func main(){
	var array = [3]int{1,2,3}
	var sum = Sum(&array)
	fmt.Println(sum)
}

func Sum(arr *[3]int)(sum int){
	for _, v:=range arr {
		sum += v
	}
	return sum
}