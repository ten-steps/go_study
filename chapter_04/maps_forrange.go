package main

import "fmt"

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	map1[5] = 5.0
	map1[6] = 6.0
	for key, val := range map1{
		fmt.Printf("map1 的key为%d，map1的值为 %f\n",key,val)
	}
}
