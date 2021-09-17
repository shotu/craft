package main

import "fmt"

// func main() {

// 	inputStr := "edef" // deeed, ddee,

// 	fmt.Println("Is isPermutationPalindrome", isPermutationPalindrome(inputStr))

// }

func isPermutationPalindrome(str string) bool {

	hashMap := map[string]int{}
	lenOfstring := len(str)
	fmt.Println("len", lenOfstring)

	for _, ele := range str {
		fmt.Println("ele", ele)
		char := fmt.Sprintf("%c", ele)
		// fmt.Println(char)
		hashMap[char]++
	}

	fmt.Println("Hashmap", hashMap)
	oddCount := 0
	for _, v := range hashMap {

		if v%2 != 0 {
			oddCount++
		}

		if oddCount > 1 {
			return false
		}
	}
	return true
}
