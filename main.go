package main

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("Son kiriting")

	var n int

	fmt.Scanf("%d", &n)

	FizzBuzz(n)
}
