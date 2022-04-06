package main

import "strings"

func reverseLeftWords(s string, n int) string {
	var sb strings.Builder

	sb.WriteString(s[n:])
	sb.WriteString(s[:n])

	return sb.String()
}
