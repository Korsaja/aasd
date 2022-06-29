package easy

// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.3 MB, less than 9.43%
func sumOfUnique(nums []int) int {
	var filter = make(map[int]int)
	var sum int
	for i := 0; i < len(nums); i++ {
		filter[nums[i]]++
	}

	for num, count := range filter {
		if count == 1 {
			sum += num
		}
	}

	return sum
}
