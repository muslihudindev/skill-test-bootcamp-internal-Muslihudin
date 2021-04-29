package main

import (
	"fmt"
	"strings"
)

func polindrome(str string) string {
	input := strings.ToLower(str)
	strTrim := strings.TrimSpace(input)
	strSplit := strings.Split(strTrim, "")

	for i := 0; i < len(strSplit)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return str + " # --> Not Palindrome"
		}
	}
	return str + " # --> Palindrome"
}

func main() {
	input1 := "Radar"
	input2 := "Malam"
	input3 := "Kasur ini rusak"
	input4 := "Ibu Ratna antar ubi"
	input5 := "Malas"
	input6 := "Makan nasi goreng"
	input7 := "Balonku ada lima"

	fmt.Println(polindrome(input1))
	fmt.Println(polindrome(input2))
	fmt.Println(polindrome(input3))
	fmt.Println(polindrome(input4))
	fmt.Println(polindrome(input5))
	fmt.Println(polindrome(input6))
	fmt.Println(polindrome(input7))

}
