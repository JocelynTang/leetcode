package golang

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	var commonPrefix string

	if len(strs) == 1 {
		commonPrefix = strs[0]
	}

	// find the shortest string index
	shortestStringIndex := 0
	for i := 1; i < len(strs); i++ {
		if len(strs[shortestStringIndex]) > len(strs[i]) {
			shortestStringIndex = i
		}
	}

	// find the common prefix
	s := strings.Split(strs[shortestStringIndex], "")
	var prefix string
	find := true

	for j := 1; j < len(s)+1; j++ {
		prefix = strings.Join(s[0:j], "")
		for _, str := range strs {
			if !(strings.HasPrefix(str, prefix)) {
				find = false
			}
		}
		if find {
			commonPrefix = prefix
			continue
		} else {
			break
		}
	}
	// save memory
	/*
		for charIndex, char := range strs[shortestStringIndex] {
			for strIndex, str := range strs {
				if strIndex == shortestStringIndex {
					continue
				}
				if rune(str[charIndex]) != char {
					return strs[shortestStringIndex][:charIndex]
				}
			}
		}
	*/

	return commonPrefix
}
