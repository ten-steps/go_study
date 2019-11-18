package main

import (
	"fmt"
	"strings"
)

type Person struct{
	firstName string
	lastName string
}


func uPerson(p *Person)  {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToLower(p.lastName)
}

func main() {
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	uPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n",pers1.firstName,pers1.lastName)

	pers2 := new(Person)
	fmt.Println(pers1)
	fmt.Println(*pers2)
	pers2.firstName = "Jake"
	pers2.lastName = "Tom"
	uPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n",pers2.firstName,pers2.lastName)

	pers3 := Person{"张三","李四"}
	fmt.Println(pers3)
	fmt.Printf("The name of the person is %s %s\n",pers3.firstName,pers3.lastName)
}