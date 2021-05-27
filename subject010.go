package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {

	//判断是否是".*"
	for i := 0; i < len(p)-1; i++ {
		if p[i] == '.' {
			if p[i+1] == '*' && i+1 == len(p)-1 {
				return true
			} else {
				return false
			}
		}
	}
	fmt.Println(".*")
	//判断是否能找到第一个字母
	start := false
	pindex := 0
	for i := 0; i < len(p); i++ {
		if p[i] == s[0] || p[i] == '.' {
			start = true
			pindex = i+1
			break
		}
	}
	if start == false {
		fmt.Println("start")
		return false
	}

	fmt.Println("end")

	if pindex == len(p) {
		if pindex == len(s) {
			return true
		} else {
			return false
		}
	}

	fmt.Println("s0", s[:1])
	fmt.Println("p0", p[:pindex])

	for sindex := 1; sindex < len(s); sindex++ {
		if s[sindex] == p[pindex] || p[pindex] == '.'{
			pindex++
		} else if p[pindex] == '*'{
			if s[sindex] != p[pindex-1] {
				return false
			}
			for p[pindex] == '*' {
				pindex++
				if pindex == len(p) {
					if pindex == len(s) {
						return true
					} else {
						return false
					}
				}
			}
		} else {
			return false
		}
		fmt.Println("sindex", sindex)
		fmt.Println("pindex", pindex)
		fmt.Println("s", s[:sindex+1])
		fmt.Println("p", p[:pindex])
		if pindex == len(p) {
			if sindex == len(s)-1 {
				return true
			} else {
				return false
			}
		}
	}
	if pindex == len(p) {
		return true
	} else {
		return false
	}
}

func main()  {
	s := "aaa"
	p := "a.a"
	fmt.Println(isMatch(s, p))
}