package main

import (
	"fmt"
	"math"
)

type SquareS struct {
	side float32
}

type Cricle struct {
	radius float32
}

type Shapers interface {
	Area() float32
}

func main() {
	var areaIntf Shapers
	sq1 := new(SquareS)
	sq1.side = 5
	areaIntf = sq1
	if t, ok := areaIntf.(*SquareS); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	switch t := areaIntf.(type) {
	case *SquareS:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Cricle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)

	}
}

func (this *SquareS) Area() float32 {
	return this.side * this.side
}

func (this *Cricle) Area() float32 {
	return this.radius * this.radius * math.Pi
}
