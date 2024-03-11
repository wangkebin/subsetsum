package main

func BFGetSingleMatchedCollection(nums []int64, targetNum int64) []int64 {
	sums := BrutalForceAllSums(nums)
	matchedIndCols := BFGetMatchedCollections(sums, len(nums), targetNum)
	if len(matchedIndCols) != 1 {
		return nil
	}
	matchedNums := make([]int64, 0)
	for _, matchedInd := range matchedIndCols[0] {
		matchedNums = append(matchedNums, nums[matchedInd])
	}
	return matchedNums
}

func BrutalForceAllSums(nums []int64) []int64 {
	res := make([]int64, 0)
	res = append(res, nums[0])
	for i := 1; i < len(nums); i++ {
		l := len(res)
		//dup := make([]int64, len(res))
		res = append(res, nums[i])
		for j := 0; j < l; j++ {
			res = append(res, nums[i]+res[j])
		}
		//res = append(res, dup...)
	}
	return res
}

func BFGetMatchedCollections(sums []int64, count int, targetNum int64) [][]int {
	ind := 0
	allres := make([][]int, 0)
	for i, v := range sums {
		if v == targetNum {
			ind = i + 1
			res := IndexToLists(ind)
			allres = append(allres, res)
		}
	}
	return allres
}

func IndexToLists(num int) []int {
	inds := make([]int, 0)
	for i := 0; true; i++ {
		if num%2 == 1 {
			inds = append(inds, i)
		}
		if num < 1 {
			break
		}
		num /= 2
	}

	return inds
}
