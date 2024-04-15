package sort

import "fmt"

func bubble(nums []int) {
	hasChanged := true
	for last := len(nums) - 1; last > 0; last-- {
		if !hasChanged {
			break
		}
		hasChanged = false
		cur := 0
		next := 1
		for ; next <= last; cur, next = cur+1, next+1 {
			if nums[cur] > nums[next] {
				nums[cur], nums[next] = nums[next], nums[cur]
				hasChanged = true
			}
		}
	}
}

func InsertionSort(nums []int) {
	for last := 1; last < len(nums); last++ {
		for cur := last; cur >= 1; cur-- {
			if nums[cur] < nums[cur-1] {
				nums[cur], nums[cur-1] = nums[cur-1], nums[cur]
			} else {
				break
			}
		}
	}
}

func mergeSort(nums []int) {

}

func Test() {
	nums := []int{2, 6, 1, 7, 9, 11, 21, 3, 5, 8, 22, 55, 3}
	bubble(nums)
	fmt.Println(nums)
}
