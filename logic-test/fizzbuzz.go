package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res

}

func main() {

	var input1 int = 3
	var input2 int = 5
	var input3 int = 15

	fmt.Println(fizzbuzz(input1))
	fmt.Println(fizzbuzz(input2))
	fmt.Println(fizzbuzz(input3))

}
