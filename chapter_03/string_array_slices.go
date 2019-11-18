package main

import "fmt"

func main() {
	s := "\u00ff\u754c"
	c := []rune(s)
	fmt.Printf("c 的数据类型为%T\n", c)
	fmt.Println(c)
	fmt.Println(string(c[:2]))

	//获取字符串的某一部分
	str := "Golang 的初学者。"
	substr := str[0 : len(str)-3]
	println(len(str))
	//println(len("。"))
	println(substr)

	//修改字符串的某个字符
	sl := []byte(str)
	sl[0] = 'p'
	println(string(sl))

	//字节数组对比函数
	res := Compare(sl, []byte(s))
	println(res)
}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	//数组长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}
