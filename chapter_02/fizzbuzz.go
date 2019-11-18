package main

const (
	Fizz     = 3
	Buzz     = 5
	FizzBuzz = 15
)

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%Fizz == 0:
			println("Fizz")
		case i%Buzz == 0:
			println("Buzz")
		case i%FizzBuzz == 0:
			println("FizzNuzz")
		default:
			println(i)
		}
	}
}
