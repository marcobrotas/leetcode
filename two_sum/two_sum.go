package two_sum

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for y := range nums {
		v := nums[y]

		if x, ok := hashMap[target-v]; ok {
			return []int{x, y}
		}

		hashMap[v] = y
	}

	return nil
}
