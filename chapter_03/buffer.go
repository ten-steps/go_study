package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	fmt.Printf("buffer type is %T\n",buffer)
	var s = "100"
	buffer.WriteString(s)
	fmt.Printf("s type is %T\n",buffer)
	//buffer.String()
	fmt.Println(buffer)
	fmt.Println(buffer.String())
}
