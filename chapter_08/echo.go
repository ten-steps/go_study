package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("help",false,"帮助信息")

const (
	Space = " "
	Newline = "\n"
)
func main() {
	//flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
