package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "abcdefg"
	// for i, v := range s {
	// 	fmt.Println(byte(v))
	// 	fmt.Println(s[i])
	// }
	l := 0
	r := 10000
	fmt.Println((l + r) >> 1)

	s := "1 5 2"
	s = strings.Trim(s, " ")

	fmt.Println(s)
}
