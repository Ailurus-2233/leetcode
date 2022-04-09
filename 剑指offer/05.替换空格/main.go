package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	var sb strings.Builder

	for _, v := range s {
		if v == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteString(string(v))
		}
	}

	return sb.String()
}

func main() {
	s := "hello world"
	fmt.Println(replaceSpace(s))
}
