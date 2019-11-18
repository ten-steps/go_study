package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	src,err := os.Open("input.dat")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer src.Close() // 函数执行完毕后关闭文件

	srcReader := bufio.NewReader(src)
	for{
		srcString,srcError := srcReader.ReadString('\n')
		if srcError == io.EOF {
			return
		}

		fmt.Printf("%s",srcString)
	}
}
