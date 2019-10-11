package main

import (
	"fmt"
	"sort"
)

// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

// 你可以假设 nums1 和 nums2 不会同时为空。

// 示例 1:

// nums1 = [1, 3]
// nums2 = [2]

// 则中位数是 2.0
// 示例 2:

// nums1 = [1, 2]
// nums2 = [3, 4]

// 则中位数是 (2 + 3)/2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	lengh := len(nums1)
	if lengh%2 != 0 {
		return float64(nums1[lengh/2])
	}
	return (float64(nums1[lengh/2]) + float64(nums1[lengh/2-1])) / 2

}

func main() {
	n1 := []int{3, 5}
	n2 := []int{2, 6}
	fmt.Printf("midle num:%f", findMedianSortedArrays(n1, n2))
}
