package main

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	result := []int{}
	xIndex := 0
	yIndex := 0
	found := false
	for x, xv := range nums {
		rest := target - xv
		b := nums[x+1:]
		for y, yv := range b {
			if rest != yv {
				continue
			}
			xIndex = x
			yIndex = x + y + 1
			found = true
		}
		if found {
			break
		}
	}

	result = append(result, xIndex, yIndex)
	return result
}
