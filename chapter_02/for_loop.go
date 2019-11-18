package main

import "fmt"

func main()  {
	foreach()
}

func foreach(){
	i := 0
	START:
		fmt.Println(i)
	i++
	if i<15 {
		goto START
	}
}