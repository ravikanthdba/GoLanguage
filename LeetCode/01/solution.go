package main

func exists(arrayIntegers []int,number int) bool {
	for num := 0; num < len(arrayIntegers); num ++ {
		if arrayIntegers[num] == number {
		return true
		}
	}
	return false
}

func twoSum(nums []int, target int) []int {
	var out []int
	for i := 0; i < len(nums) - 1; i ++ {
		for j := 0; j < len(nums); j ++ {
			if i != j {
				if nums[i] + nums[j] == target {
					if !exists(out, i) || !exists(out, j) {
					out = append(out, i)
					out = append(out, j)
					}
				}
			}
		}
	}
	return out
}