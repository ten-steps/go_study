package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var orig string = "ABC"
	var newS string
	fmt.Printf("The size ints is:%d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is Integer - exiting with error\n",err)
		os.Exit(1)
	}
	fmt.Printf("The integer is %d\n",an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is %s\n",newS)
}
