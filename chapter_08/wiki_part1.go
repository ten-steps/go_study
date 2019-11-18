package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	wiki := Page{"test.txt",[]byte("hello wrold!")}
	wiki.save()
}

func (this *Page)save(){
	f,err := os.OpenFile(this.Title,os.O_CREATE|os.O_WRONLY,0)
	if err != nil {
		fmt.Println("文件打开或创建失败")
	}
	f.WriteString(string(this.Body))
}