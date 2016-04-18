package main

import (
	"fmt"
	"strconv"
)

// Program that prints out 1 to 100. Everytime a number is multiple with 3, it prints "fizz"
// everytime the number is multible with 5, it prints "buzz"
// everytime a number is mutible with 3 & 5, it prints "fizzbuzz"
func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

func FizzBuzz(i int) string {
	if i%3 == 0 && i%5 == 0 { // checks if "i" is multiple with 3 & 5
		return "fizzbuzz"
	} else if i%3 == 0 { // checks if "i" is multiple with 3
		return "fizz"
	} else if i%5 == 0 { // checks if "i" is multiple with 5
		return "buzz"
	} else {
		return strconv.Itoa(i) // if not multiple with 3 or 5, prints the number
	}
}
