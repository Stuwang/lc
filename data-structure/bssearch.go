package data_structure

import "log"

func search(nums []int, f int) int {
	s := 0
	e := len(nums)
	for s < e {
		mid := s + (e-s)/2
		n := nums[mid]
		if n == f {
			return mid
		} else if n < f {
			s = mid + 1
		} else {
			e = mid
		}
	}
	return -1
}

func testS() {
	log.Println(search([]int{1,2,4}, 3))
}