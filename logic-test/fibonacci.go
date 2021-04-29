package main

import "fmt"

func fibonacci(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}

func fibonacciRange(data []int) int {
	var total int
	for _, v := range data {
		total += v
	}

	var i int
	var fibo int

	for total > fibonacci(i) {
		fibo = fibonacci(i + 1)
		i++
	}

	div := fibo - total

	return div
}

func main() {
	data := []int{15, 1, 3}
	fmt.Println(fibonacciRange(data))
}
