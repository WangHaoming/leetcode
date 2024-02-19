package towsum

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}

func Test() {
	nums := []int{3, 7, 4, 5}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
