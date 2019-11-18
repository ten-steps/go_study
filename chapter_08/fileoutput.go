package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile,outputError := os.OpenFile("output.dat",os.O_WRONLY|os.O_CREATE,0666);
	if outputError != nil {
		fmt.Println("文件打开或创建时出错")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i:=0;i<100 ;i++  {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
