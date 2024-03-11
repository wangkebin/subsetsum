package main

import (
	"sort"
	"time"
)

func BranchBound(nums []int64, targetNum int64) []int64 {
	selected := make([]int64, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	sumRest := make([]int64, len(nums))
	sumr := int64(0)
	for j := len(nums) - 1; j >= 0; j-- {
		sumr += nums[j]
		sumRest[j] = sumr
	}
	c := make(chan []int64)
	go func() {
		BBChan(nums, sumRest, targetNum, 0, 0, selected, c)
	}()

	select {
	case selected := <-c:
		return selected
	case <-time.After(30 * time.Second):
		return nil
	}
}

func BBChan(nums []int64, sumRest []int64, targetNum int64, ind int, sum int64, selected []int64, c chan []int64) {
	c <- BranchBoundRecursive(nums, sumRest, targetNum, 0, 0, selected)
}

func BranchBoundRecursive(nums []int64, sumRest []int64, targetNum int64, ind int, sum int64, selected []int64) []int64 {
	if ind == len(nums) {
		return nil
	}

	// upper bound
	if sum+sumRest[ind] < targetNum {
		return nil
	}

	// including current nums[ind]
	nextSum := sum + nums[ind]
	if nextSum == targetNum {
		return append(selected, nums[ind])
	} else if nextSum < targetNum {
		rest := BranchBoundRecursive(nums, sumRest, targetNum, ind+1, nextSum, append(selected, nums[ind]))
		if len(rest) > 1 {
			return rest
		}
	}

	// NOT including current nums[ind]
	rest := BranchBoundRecursive(nums, sumRest, targetNum, ind+1, sum, selected)
	if len(rest) > 1 {
		return rest
	}

	// nil when no match
	return nil
}
