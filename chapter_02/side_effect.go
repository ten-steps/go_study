package main

func main() {
	n := 0
	reply := &n
	Multiply(5,4,reply)
	println(n)
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
