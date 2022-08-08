package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	sort.Slice(nums, func(i, j int) bool {
		a := fmt.Sprintf("%d%d", nums[i], nums[j])
		b := fmt.Sprintf("%d%d", nums[j], nums[i])
		return strings.Compare(a, b) < 0
	})
	s := ""
	for _, v := range nums {
		s += strconv.Itoa(v)
	}
	return s
}
