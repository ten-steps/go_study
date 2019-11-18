package main

import "fmt"

type Address struct {
	province string
	city string
	county string
	street string
}

type Vcard struct {
	name string
	phone int
	brith string
	avatar string
	address Address
}

func main() {
	vcard := new(Vcard)
	address := new(Address)
	address.province = "河南省"
	address.city = "郑州市"
	address.county = "经八街道"
	address.street = "纬一路9号家属院"
	vcard.name =  "张冉"
	vcard.phone = 17796701097
	vcard.brith = "1998.12.23"
	vcard.avatar = "public/01.png"
	vcard.address = *address
	fmt.Println(*vcard)
}
