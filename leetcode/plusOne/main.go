package plusOne

import "fmt"

func plusOne(nums []int) []int {
	last := len(nums) - 1
	carry := 0
	nums[last] = nums[last] + 1
	for last >= 0 {
		currentBit := nums[last] + carry
		if currentBit == 10 {
			carry = 1
			currentBit = 0
		} else {
			carry = 0
		}
		nums[last] = currentBit
		last--
	}

	if carry == 1 {
		nums = append([]int{1}, nums...)
	}

	return nums
}

func plusOne2(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i]++
		if nums[i] < 10 {
			return nums
		}
		nums[i] = 0
	}

	return append([]int{1}, nums...)

}

func Test() {
	nums := []int{1, 2, 3, 4, 5, 7, 8, 9}
	// nums := []int{9, 9, 9}
	// nums := []int{1, 3, 4}
	result := plusOne2(nums)
	fmt.Println(result)
}
