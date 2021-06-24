package main

import "fmt"

func replacespace(s string) string {
	b := []byte(s)
	len_s := len(b)
	space_num := 0
	for _, v := range b {
		if v == ' ' {
			space_num++
		}
	}
	temp := make([]byte, 2*space_num)
	b = append(b, temp...)
	len_b := len(b)
	j := len_s - 1
	for i := len_b - 1; i >= 0; i-- {
		if b[j] == ' ' {
			j--
			b[i] = '0'
			b[i-1] = '2'
			b[i-2] = '%'
			i -= 2
		} else {
			b[i] = b[j]
			j--
		}
	}
	return string(b)
}

func main() {
	s := "  Who are you ? "
	res := replacespace(s)
	fmt.Println(res)
}
