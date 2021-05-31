package main

import "fmt"

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		fmt.Println(len(s1), len(s2))
		fmt.Println(s1[0]-'a', s2[0]-'a')
		return false
	}
	exists := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		if v, ok := exists[s1[i]]; ok {
			exists[s1[i]] = v + 1
		} else {
			exists[s1[i]] = 1
		}
		fmt.Println(string(s1[i]), exists[s1[i]])
	}
	for i := 0; i < len(s2); i++ {
		if v, ok := exists[s2[i]]; ok && v > 0 {
			exists[s2[i]]--
			fmt.Println(string(s2[i]), exists[s2[i]])
		} else {
			return false
		}
	}
	return true
}

func main() {
	s1 := "xiaoming"
	s2 := "mingxiao"
	tt := isAnagram(s1, s2)
	fmt.Println(tt)
}
