package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxlen := 0
	for start := 0; start < len(s)-maxlen; start++ {
		for end := start+maxlen+1; end <= len(s); end ++ {
			d := s[start:end]
			if norepeatletter(d) {
				maxlen = len(d)
			} else  {
				break
			}
		}
	}
	return maxlen
}

func norepeatletter(s string) bool {
	substring := make(map[int32]int)
	for _,v := range s {
		if _, ok := substring[v]; ok {
			substring[v]++
			if substring[v]>1 {
				return false
			}
		} else {
			substring[v] = 1
		}
	}
	return true
}

func main() {
	s := "psdlscfkjsdf"
	fmt.Println(lengthOfLongestSubstring(s))
}


