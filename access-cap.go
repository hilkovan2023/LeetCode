package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	N := []int{10,20,30,40,50,60,70,80,90,100, 110,120,130,140,150,160,170,180,190,200,210,220,230,240,250,260,270,280,290,300} //传感器数量
	R := 10  //重复试验次数
	//chongfu := 10  //重复试验次数
	tongguo := [][]int{}
	res := []int{}  //传感器发送时间
	m := []int{}  //冲突传感器列表

	for i:=0; i<len(N); i++ {
		for j:=0; j<R; j++ {
			rand.Seed(time.Now().Unix())
			fmt.Println("N:", N[i])

			res = res[0:0]
			for p:=0; p<N[i]; p++ {
				res = append(res, rand.Intn(2000))
			}
			fmt.Println(res)
			sort.Ints(res)
			fmt.Println(res)

			m = m[0:0]
			for p:=1; p<N[i]-1; p++ {
				if res[p+1]-res[p] < 25 || res[p]-res[p-1] < 25{
					m = append(m, res[p])
				}
			}
			fmt.Println(m)
			fmt.Println("len:", len(m))
			tongguo = append(tongguo, []int{})
			tongguo[i] = append(tongguo[i], N[i]-len(m)-2)
		}
	}
	fmt.Println("tongguo:", tongguo)

}
