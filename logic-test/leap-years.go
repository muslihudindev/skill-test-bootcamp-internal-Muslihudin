package main

import (
	"fmt"
	"strings"
)

func leapyears(start int, end int) []int {

	var res []int
	for i := start; i <= end; i++ {
		if i%4 == 0 {
			if i%100 == 0 {
				if i%400 == 0 {
					res = append(res, i)
				}
			} else {
				res = append(res, i)
			}
		}
	}
	return res

}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func main() {
	var start int = 1900
	var end int = 2020

	data := leapyears(start, end)
	fmt.Println(arrayToString(data, ", "))
}
