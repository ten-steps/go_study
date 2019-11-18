package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func main() {
	nrchars, nrwords, nrlines = 0, 0, 0
	inpresReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some inpres, type S to stop: ")
	for {
		inpres, err := inpresReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		if inpres == "S\r\n" { // Windows, on Linux it is "S\n"
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Counters(inpres)
	}
}

func Counters(inpres string) {
	nrchars += len(inpres) - 2 // -2 for \r\n
	nrwords += len(strings.Fields(inpres))
	nrlines++
}