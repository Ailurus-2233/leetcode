package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isNumber("96 e5"))
}

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		arr := strings.Split(s, "e")
		if len(arr) == 2 {
			return isFloat(arr[0]) && isInt(arr[1])
		}
		arr = strings.Split(s, "E")
		if len(arr) == 2 {
			return isFloat(arr[0]) && isInt(arr[1])
		}
		return false
	} else {
		return true
	}
}

func isInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func isFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
