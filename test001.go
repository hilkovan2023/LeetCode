/*此处为答题区域，请直接将代码写在下方*/
package main

import "fmt"

//type status string

func statuschange(table map[string]string, first_status string, actions []string) []string{
	res := []string{}
	cur_status := first_status
	for i:=0; i<len(actions); i++ {
		if _, ok := table[cur_status+actions[i]]; !ok {
			res = append(res, "[no action!]"+actions[i])
		} else {
			cur_status = table[cur_status+actions[i]]
			res = append(res, cur_status)
		}
	}
	return res
}

//func (s status) event(table map[status]string, action string) string {
//	return table[s]
//}

func main() {
	status := []string{"Opened", "Closed", "Unlocked", "Locked"}
	action := []string{"closeDoor", "openDoor", "lockDoor", "unlockDoor"}

	table := map[string]string {
		status[0]+action[0]:status[1],
		status[1]+action[1]:status[0],
		status[1]+action[2]:status[3],
		status[3]+action[3]:status[2],
		status[2]+action[1]:status[0],
		status[2]+action[2]:status[3],
	}
	actions := []string{"closeDoor", "lockDoor", "unlockDoor", "openDoor"}
	first_status := "Opened"
	results := []string{"Closed", "Locked", "Unlocked", "Opened"}
	res := statuschange(table, first_status, actions)

	fmt.Println(res)
	for i:=0; i<len(actions); i++ {
		if results[i] == res[i] {
			fmt.Println(results[i], res[i], "True")
		} else {
			fmt.Println(results[i], res[i], "False")
		}
	}
}