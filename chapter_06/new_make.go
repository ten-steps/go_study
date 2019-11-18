package main

type bar struct {
	thingOne string
	thingTwo int
}

func main() {
	bar := NewBar("张三",18)
}

func NewBar(string2 string,int2 int)*bar  {
	m := new(bar)
	m.thingOne = string2
	m.thingTwo = int2
	return m
}