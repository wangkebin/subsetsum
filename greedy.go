package main

import "sort"

func Greedy(nums []int64, targetNum int64) []int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < len(nums); i++ {
		res := GreedyMatching(nums[i:], targetNum)
		if res != nil {
			return res
		}
	}
	return nil

}

func GreedyMatching(nums []int64, targetNum int64) []int64 {

	matched := make([]int64, 0)
	for _, v := range nums {
		if targetNum <= 0 {
			break
		}
		if targetNum >= v {
			matched = append(matched, v)
			targetNum -= v
		}
	}
	if targetNum == 0 {
		return matched
	}
	return nil
}
