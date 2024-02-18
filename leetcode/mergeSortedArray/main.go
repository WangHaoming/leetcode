package mergeSortedArray

import "fmt"

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {

	idx1, idx2 := m-1, n-1
	last := len(nums1) - 1
	for last >= 0 {
		for idx2 >= 0 && idx1 >= 0 {
			if nums2[idx2] >= nums1[idx1] {
				nums1[last] = nums2[idx2]
				idx2--
				last--
				break
			} else {
				nums1[last] = nums1[idx1]
				idx1--
				last--
				break
			}
		}
		if idx2 < 0 {
			break
		}

		if idx1 < 0 {
			nums1[last] = nums2[idx2]
			idx2--
			last--
		}

	}

}

func Test() {
	nums1 := []int{7, 8, 9, 9, 11, 0, 0, 0}
	nums2 := []int{4, 6, 9}
	mergeSortedArray(nums1, 5, nums2, 3)
	fmt.Println(nums1)
}
