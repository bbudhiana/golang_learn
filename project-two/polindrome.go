package main

import "fmt"

func main() {
	fmt.Println("Polindrome Testing")

	fmt.Println(isPolindrome("lkjjkl"))
}

func isPolindrome(word string) bool {

	for i := 0; i < len(word)/2; i++ {
		//fmt.Println(i, "is", string(word[i]))
		//fmt.Println(i, "again", len(word)-i-1)

		index1 := i
		index2 := len(word) - i - 1
		if word[index1] != word[index2] {
			return false
		}

	}
	return true
}
