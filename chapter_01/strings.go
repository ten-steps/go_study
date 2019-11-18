package main

import (
	"fmt"
	"strings"
)

func main()  {
	var s = "pre_qwasqzxerqMNdfcqv.txt"

	//前缀 和 后缀
	var pre = strings.HasPrefix(s,"pre_")
	var ext = strings.HasSuffix(s,".txt")
	println(pre,ext)

	//字符串包含关系
	println(strings.Contains(s,"er"))

	//判断字符串在父字符串首次出现的位置
	var sp = strings.Index(s,"1")
	var sp1 = strings.Index(s,"q")
	var lastsp = strings.LastIndex(s,"q")
	println(sp,sp1,lastsp)

	a := 'a'
	var spRune = strings.IndexRune(s,a)
	println(spRune)

	//替换字符串
	var news = strings.Replace(s,"q","e",2)
	//n是替换的数目，n < 0 时会替换所有

	//重复字符串
	var reply = strings.Repeat(s,5)
	println(news,reply)

	//修改字符串大小写
	var  toUpperStr  = strings.ToUpper(s)
	var toLowerStr = strings.ToLower(s)
	println(toUpperStr,toLowerStr)

	//修剪字符串
	var bad = "  kjj;skjs'fkd;gjkdhogh   "
	println(bad,strings.TrimSpace(bad))

	//分割字符串
	segmentation := "张三、李四、王二、赵六"
	newSeg :=  strings.Split(segmentation,"、")
	fmt.Println(newSeg)
}