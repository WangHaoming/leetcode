package removeDuplicates

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

func rmDuplicate(nums []int) int {
	last := len(nums) - 1
	current := 0
	for current < last {
		first := 0
		for first < current {
			if nums[current] == nums[first] {
				tmp := nums[current]
				nums[current] = nums[last]
				nums[last] = tmp
				last--
			} else {
				first++
			}
		}
		current++
	}
	if current == last {
		return current + 1
	} else {
		return current
	}
}

func Test() {
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	nums := []int{1, 2, 3, 4, 5, 7, 8, 9}

	// result := rmDuplicate(nums)
	result := removeDuplicates(nums)
	fmt.Println(result)
}
