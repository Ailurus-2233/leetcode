package main

import "fmt"

// func 1
//func hammingWeight(num uint32) int {
//	ans := 0
//	for {
//		if num == 0 {
//			break
//		}
//		if num&1 == 1 {
//			ans++
//		}
//		num = num >> 1
//	}
//	return ans
//}

// func 2
func hammingWeight(num uint32) int {
	ans := 0
	for {
		if num == 0 {
			break
		}
		ans++
		num = num & (num - 1)
	}
	return ans
}

func main() {
	num := uint32(0b00000000000000000000000000001011)
	fmt.Println(hammingWeight(num))
}
