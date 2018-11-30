package main

import "fmt"

func fibonacci() func(int) int {
	return func(i int) int{
		if i > 1{
			return (fibonacci()(i-1) + fibonacci()(i-2))
		}
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}