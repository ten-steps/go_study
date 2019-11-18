package main

import "fmt"

func main() {
	var value int
	var isPresent bool
	map1 := make(map[string]int)
	map1["beijing"] = 50
	map1["shanghai"] = 65
	map1["郑州"] = 101
	value,isPresent = map1["郑州"]
	if isPresent {
		fmt.Printf("The value of '郑州' in map1 is：%d\n",value)
	}else{
		fmt.Printf("map1 does not contain 郑州")
	}
	delete(map1,"beijing")
	value,isPresent = map1["beijing"]
	if isPresent {
		fmt.Printf("The value of 'beijing' in map1 is：%d\n",value)
	}else{
		fmt.Printf("map1 does not contain beijing")
	}
	fmt.Println(map1)
}
