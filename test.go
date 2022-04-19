package main

import "fmt"

func main() {
	s := "abcdefg"
	for i, v := range s {
		fmt.Println(byte(v))
		fmt.Println(s[i])
	}
}
