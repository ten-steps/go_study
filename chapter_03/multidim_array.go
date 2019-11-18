package main

import "fmt"

const (
	WIDTH  = 5
	HEIGHT = 3
)

type pixel int

var screen [WIDTH][HEIGHT] pixel

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	fmt.Println(*&screen)
	s := screen[2:5] //多维切片
	fmt.Println(s)
}
