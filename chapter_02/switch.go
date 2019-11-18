package main

import "fmt"

func main()  {
	k:=6
	switch k {
	case 4: fmt.Println("was <= 4")
	case 5: fmt.Println("was <= 5")
	case 6: fmt.Println("was <= 6")
	case 7: fmt.Println("was <= 7")
	case 8: fmt.Println("was <= 8")
	default: fmt.Println("default case")
	}
}
