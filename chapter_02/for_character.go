package main

import "strings"

func main() {
	forChacter()
}

func forChacter() {
	//for i := 0; i < 25; i++ {
	//	for j := 0; j < i; j++ {
	//		print("G")
	//	}
	//	println()
	//}
	for i := 0; i < 25; i++ {
		println(strings.Repeat("G",i))
	}
}
