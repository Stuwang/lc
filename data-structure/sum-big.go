package data_structure

import "fmt"

func sumBig(nums1, nums2 []int64) (res []int64) {
	l := maxInt(len(nums1), len(nums2)) + 1
	nums1 = resAndAppend(nums1, l)
	nums2 = resAndAppend(nums2, l)
	res = make([]int64, l)
	var m int64
	for i := 0; i < l; i++ {
		res[i] = m + nums1[i] + nums2[i]
		m = res[i] / 10
		res[i] %= 10
	}
	for i := 0; i < len(res)/2; i++ {
		t := res[i]
		res[i] = res[len(res)-i-1]
		res[len(res)-i-1] = t
	}
	return res
}

func maxInt(l, r int) int {
	if l < r {
		return r
	}
	return l
}

func resAndAppend(nums []int64, l int) []int64 {
	for i := 0; i < len(nums)/2; i++ {
		t := nums[i]
		nums[i] = nums[len(nums)-i-1]
		nums[len(nums)-i-1] = t
	}
	res := nums
	for i := 0; i < l-len(nums); i++ {
		res = append(res, 0)
	}
	return res
}

func RunSumBig() {
	fmt.Println(sumBig([]int64{1, 2, 3}, []int64{1, 2}))
}
