package main

import (
	"fmt"
	"reflect"
)

type Personer struct {
	name string "这是一个人的姓名"
	sex string  "性别"
	age int  "年龄"
}
func main() {
	tt := new(Personer)
	tt.name = "张三"
	tt.sex = "男"
	tt.age = 18
	for i:= 0; i < 3 ;i++  {
		refTag(*tt,i)
	}
}

func refTag(tt Personer,ix int){
	ttType := reflect.TypeOf(tt)
	ixField :=  ttType.Field(ix)
	fmt.Printf("%v\n",ixField.Tag)
}