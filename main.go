package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . num num num")
		fmt.Println("Example: go run . 2 5 4")
		return
	}
	ags := os.Args[1:]
	numSlice := []int{}

	for _, str := range ags {
		for _, n := range str {
			if n < '0' && n > '9' {
				fmt.Println("Enter numbers only")
				return
			}
		}
	}

	for _, nb := range ags {
		num, _ := strconv.Atoi(nb)
		numSlice = append(numSlice, num)
	}
	fmt.Println(BetweenUs(numSlice))
}

// Function BetweenUs() take a slice of 3 numbers and checks
// if the first number is between the secon and the third, then returns a bool.
func BetweenUs(s []int) bool {
	if s[0] >= s[1] && s[0] < s[2] {
		return true
	}
	return false
}
