package main

import "fmt"

func main() {
	FizzBuzz(100)
}

func FizzBuzz(loopCount int) {
	for i := 1; i <= loopCount; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		} else if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
