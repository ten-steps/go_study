package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
)

func main(){
	for i := 0;i < 10;i++{
		a:=rand.Int()
		fmt.Printf("%d /",a)
	}
	print("\n")
	for i := 0;i<5;i++{
		r := rand.Intn(8)
		fmt.Printf("%d /",r)
	}
	print("\n")
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i:= 0;i<10;i++{
		fmt.Printf("%2.2f /" ,100*rand.Float32())
	}
}