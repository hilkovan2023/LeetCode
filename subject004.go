package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if nums1 == nil {
		if len(nums2)%2 == 1 {
			return float64(nums2[len(nums2)/2])
		} else {
			return float64((nums2[len(nums2)/2] + nums2[len(nums2)/2-1])) / 2
		}
	} else if nums2 == nil {
		if len(nums1)%2 == 1 {
			return float64(nums1[len(nums1)/2])
		} else {
			return float64((nums1[len(nums1)/2] + nums1[len(nums1)/2-1])) / 2
		}
	} else {
		nums3 := sortt(nums1, nums2)
		fmt.Println(nums3)
		if len(nums3)%2 == 1 {
			return float64(nums3[len(nums3)/2])
		} else {
			return float64((nums3[len(nums3)/2] + nums3[len(nums3)/2-1])) / 2
		}
	}
}

//两个正序数组拼接成一个正序数组
func sortt(nums1 []int, nums2 []int) []int {
	nums3 := make([]int, len(nums1)+len(nums2))
	index1 := 0
	index2 := 0
	index3 := 0
	for index1 < len(nums1) || index2 < len(nums2) {
		if index1 == len(nums1) {
			nums3[index3] = nums2[index2]
			index2++
			index3++
		} else if index2 == len(nums2) {
			nums3[index3] = nums1[index1]
			index1++
			index3++
		} else {
			if nums1[index1] <= nums2[index2] {
				nums3[index3] = nums1[index1]
				index1++
				index3++
			} else {
				nums3[index3] = nums2[index2]
				index2++
				index3++
			}
		}
	}
	return nums3
}

func main() {
	nums1 := []int{}
	nums2 := []int{3, 4, 5, 8}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
