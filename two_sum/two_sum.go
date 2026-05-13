package main

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	len_nums := len(nums)
	for i := 0; i < len_nums; i++ {
		for j := i + 1; j < len_nums; j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				break
			}
		}
	}
	return result
}
