// p(m)数组和q(n)数组
// p中找包含q的最小子串
// 复杂度 O(m*m*n)
// 进一步降低复杂度 O(m+n)

package main

import (
	"fmt"
)

func searchminstr(p []byte, q []byte) []byte {
	if len(p) < len(q) {
		return nil
	}
	if !exitstr(p, q) {
		return nil
	}
	//matchnum := 0
	mfast := len(p)
	mslow := 0
	slow := 0
	fast := len(q)
	//qmap := make()
	for ; fast <= len(p); fast++ {
		fmt.Println("for1", string(p[slow:fast]))
		for exitstr(p[slow:fast], q) {
			fmt.Println("for2", string(p[slow:fast]))
			if fast-slow < mfast-mslow {
				mfast = fast
				mslow = slow
			}
			slow++
		}
	}
	return p[mslow:mfast]
}

func exitstr(p []byte, q []byte) bool {
	flag := false
	for i:=0; i<len(q); i++ {
		flag = false
		for j:=0; j<len(p); j++ {
			if q[i] == p[j] {
				flag = true
				//fmt.Println("q[",i,"]","=","p[",j,"]")
				break
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}


func main() {
	p := []byte("absdersf")
	q := []byte("sf")
	//fmt.Println(reflect.DeepEqual(p[0:3], q[0:3]))
	fmt.Println("exitstr(p, q):", exitstr(p, q))
	fmt.Println(string(searchminstr(p, q)))
}

