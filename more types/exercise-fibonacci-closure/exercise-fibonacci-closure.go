package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
var (
	a = 0
	b = 1
	c = 0
)

func fibonacci() func() int {
	return func() int {
		c, a, b = a, b, a+b
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
