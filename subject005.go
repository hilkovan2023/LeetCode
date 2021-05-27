package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	} else if len(s) == 1 {
		return s
	} else if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}

	maxcenter := 0
	maxwidth := 0
	for center := 0; center < len(s)-maxwidth; center++ {
		for width := 0; center-width >= 0 && len(s)-center > width; width++ {
			if s[center-width] == s[center+width] {
				fmt.Println(center, width)
				if maxwidth < width {
					maxwidth = width
					maxcenter = center
					fmt.Println("sdf", maxwidth, maxcenter)
				}
			} else {
				break
			}
		}
	}
	flag := 0
	for center := maxwidth; center < len(s)-maxwidth-1; center++ {
		for width := 0; center-width>=0 && len(s)-center-1>width; width++ {
			if s[center-width] == s[center+1+width] {
				if maxwidth <= width {
					flag = 1
					maxwidth = width
					maxcenter = center
					fmt.Println("dddd", maxcenter, maxwidth)
				}
			} else {
				break
			}
		}
	}
	if flag == 1 {
		return string(s[maxcenter-maxwidth: maxcenter+1+maxwidth+1])
	} else {
		return string(s[maxcenter-maxwidth: maxcenter+maxwidth+1])
	}
}

func main() {
	s := "111111111111"
	fmt.Println(len(longestPalindrome(s)), longestPalindrome(s))
}
