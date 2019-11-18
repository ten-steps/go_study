package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	iReader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Please enter your name： ===")
	i,err :=  iReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading,exiting program.")
		return
	}

	fmt.Printf("Your name  is %s",i)

	switch i{
	case "Philip\r\n":  fmt.Println("Welcome Philip!")
	case "Chris\r\n":   fmt.Println("Welcome Chris!")
	case "Ivo\r\n":     fmt.Println("Welcome Ivo!")
	default: fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch i {
	case "Philip\r\n":  fallthrough
	case "Ivo\r\n":     fallthrough
	case "Chris\r\n":   fmt.Printf("Welcome %s\n", i)
	default: fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch i {
	case "Philip\r\n", "Ivo\r\n":   fmt.Printf("Welcome %s\n", i)
	default: fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}
