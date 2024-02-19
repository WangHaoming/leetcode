package searchInsertPosition

import "fmt"

func searchInsertPosition(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

func searchInsertPositionLeecode(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}

	}
	return 0
}

func searchInsertPosition2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	last := (len(nums) - start) / 2
	for start < last {
		if start == len(nums)-1 {
			break
		}
		step := func() int {
			if (last-start)/2 == 0 {
				return 1
			}
			return (last - start) / 2
		}()
		if target == nums[last] {
			return last
		}
		if target < nums[last] {
			last = last - step
		}
		if target > nums[last] {
			tmp := last
			last = last + step
			start = tmp
		}
	}
	return start
}

func searchInsertPosition3(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if target <= nums[low] {
		return low
	} else {
		return low + 1
	}
}

func Test() {
	nums := []int{2, 5, 6, 9, 11}
	targets := 12
	result := searchInsertPositionLeecode(nums, targets)
	fmt.Println(result) // 输出结果，例如[0, 1]
}
