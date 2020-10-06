package main

import (
	"fmt"
)

func main() {
	fmt.Println(validPalindrome("eccer"))
}

func validPalindrome(s string) bool {

	lenStr := len(s)

	if lenStr == 1 {
		return true
	}

	if lenStr == 2 {
		return true
	}

	isPalindrome, failIndex := checkPalindrome(s)

	if isPalindrome {
		return true
	}

	str1 := make([]byte, lenStr-1)
	str2 := make([]byte, lenStr-1)
	j, k := 0, 0

	for i := range s {
		if i != failIndex {
			str1[k] = s[i]
			k++
		}

		if i != len(s)-1-failIndex {
			str2[j] = s[i]
			j++
		}
	}

	p1, _ := checkPalindrome(string(str1))
	p2, _ := checkPalindrome(string(str2))

	return p1 || p2
}

func checkPalindrome(s string) (bool, int) {

	halfLen := len(s) / 2

	for i := 0; i < halfLen; i++ {
		if s[i] != s[len(s)-1-i] {
			return false, i
		}
	}

	return true, halfLen
}