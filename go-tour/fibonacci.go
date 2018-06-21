package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f := 0
	return func() int {
		first := 0
		second := 1
		for i := 0; i < f; i++ {
			first, second = second, first+second
		}
		f++
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
