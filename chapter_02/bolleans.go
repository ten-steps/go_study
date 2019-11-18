package main

import (
	"fmt"
	"runtime"
)
var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."
func init()  {
	if runtime.GOOS == "windows" {
		prompt =fmt.Sprintf(prompt,"Ctrl+Z,Enter")
	}else{
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}
func main(){
	bool1 := true
	if bool1 {
		fmt.Printf("the value is true\n")
	}else{
		fmt.Printf("This value is false \n")
	}

	system := runtime.GOOS
	println(system)
}
