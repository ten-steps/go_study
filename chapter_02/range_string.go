package main

import "fmt"

func main(){
	//var str string = "chinese : Go 语言入门体验"
	//for pos,val:=range str{
	//	fmt.Printf("%c,%d\n",val,pos)
	//}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
