package main

func main() {
	str := "abcdefj"
	var str1, str2 = cutStr(str, 5)
	println(str1, str2)

	str3 := "张三"
	newstr := strtev(str3)
	news := swopStr(str)
	println(newstr,news)
}

//接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。

func cutStr(str string, i int) (nstr1, nstr2 string) {
	l := len(str)
	if i > l {
		nstr1 = str
		return nstr1, nstr2
	} else {
		strl := []byte(str)
		nstr1 = string(strl[0:i])
		nstr2 = string(str[i:])
		return nstr1, nstr2
	}
}

//反转字符串
func strtev(str string) string {
	s := []rune(str)
	len := len(s)
	cont := len-1
	v := make([]rune, len)
	for i := 0; i <= cont; i++ {
		v[cont-i] = s[i]
	}
	return string(v)
}

func swopStr(str string) string{
	s := []rune(str)
	strLen := len(s)
	for i:=0;i < strLen/2; i++ {
		str1 := s[i]
		s[i] = s[strLen-1-i]
		s[strLen-1-i] = str1
	}
	return string(s)
}