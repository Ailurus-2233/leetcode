package main

import (
	"fmt"
)

func firstUniqChar(s string) byte {
	data := [26]int{}
	for _, v := range s {
		data[v-'a'] += 1
	}
	for _, v := range s {
		if data[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}

func main() {
	s := "abaccdeff"
	fmt.Println(firstUniqChar(s))
}
