package main

import "strings"

func lengthOfLastWord(s string) int {
	res := strings.Fields(s)
	return len(res[len(res)-1])
}
