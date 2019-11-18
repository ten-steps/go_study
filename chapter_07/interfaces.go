package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}


func main() {
	sq1 := new(Square)
	sq1.side =5
	area := sq1
	fmt.Printf("The square has area: %f\n", area.Area())
}

func (this Square) Area() float32  {
	return this.side * this.side
}