package main

import "fmt"

func quickSort(nums []int, left, right int) []int {
	if left > right {
		return nil
	}

	i, j, point := left, right, nums[left]
	for i < j {
		for i < j && nums[j] >= point {
			j--
		}
		for i < j && nums[i] <= point {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort(nums, left, j-1)
	quickSort(nums, j+1, right)
	return nums
}
func sort2(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}

func main() {
	nums := []int{1, 3, 6, 2, 5, 4, 8, 2, 6}
	sort2(nums)
	fmt.Println(nums)
}
