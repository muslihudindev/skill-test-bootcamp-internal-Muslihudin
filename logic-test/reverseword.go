package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	kata := "I am A Great human"
	fmt.Println(kata)
	fmt.Println(" into ")
	fmt.Println(wordreverse(kata))
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func wordreverse(kata string) string {
	arrkata := strings.Split(kata, " ")
	var arrres []string
	for _, v := range arrkata {
		var res string
		con := strings.Split(v, "")
		stat := false
		for i := len(con) - 1; i >= 0; i-- {
			res += con[i]
			if IsUpper(con[i]) {
				stat = true
			}
		}
		if stat {
			res = strings.Title(strings.ToLower(res))
		}
		arrres = append(arrres, res)
	}

	result := strings.Join(arrres, " ")
	return result
}
