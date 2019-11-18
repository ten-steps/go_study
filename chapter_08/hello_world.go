package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name += strings.Join(os.Args[1:]," ")
	}
	fmt.Println("hello wrold",name)
}
