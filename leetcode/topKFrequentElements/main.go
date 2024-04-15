package topkfrequentelements

import (
	// "container/heap"
	"fmt"
)

func getParentIdx(cur int) int {
	if cur == 0 {
		return -1
	}
	return (cur - 1) / 2
}

func getChildrenIdx(cur int, len int) (int, int) {

	cidx1 := cur*2 + 1
	if cidx1 >= len {
		return -1, -1
	}
	cidx2 := cidx1 + 1
	if cidx2 >= len {
		return cidx1, -1
	}
	return cidx1, cidx2
}

type BiHeap struct {
	pq PriorityQueue
}

func (bh *BiHeap) Push(i Item) {
	pq := &((*bh).pq)
	(*pq).Push(i)
	cur := pq.Len() - 1
	bh.siftUp(cur)

	// too complicated for using for doing siftup
	// p := getParentIdx(cur)
	// for p >= 0 {
	// 	if pq.Less(cur, p) {
	// 		pq.Swap(cur, p)
	// 	} else {
	// 		break
	// 	}
	// 	cur = p
	// 	p = getParentIdx(cur)
	// }
}

func (bh *BiHeap) siftUp(cur int) {
	pq := &((*bh).pq)
	if cur == 0 {
		return
	}
	p := getParentIdx(cur)
	if pq.Less(cur, p) {
		pq.Swap(cur, p)
		bh.siftUp(p)
	} else {
		return
	}
}

func (bh *BiHeap) siftDown(cur int) {
	pq := &((*bh).pq)
	cidx1, cidx2 := getChildrenIdx(cur, pq.Len())
	selected := bh.getBigestChild(cidx1, cidx2)
	if selected < 0 {
		return
	} else {
		if !pq.Less(cur, selected) {
			pq.Swap(cur, selected)
			bh.siftDown(selected)
		}
	}

}

func (bh *BiHeap) getBigestChild(left, right int) int {
	pq := &((*bh).pq)
	if left < 0 {
		return -1
	} else if right < 0 {
		return left
	}
	if pq.Less(left, right) {
		return left
	} else {
		return right
	}
}

func (bh *BiHeap) Pop() Item {
	pq := &((*bh).pq)
	item := (*pq)[0]
	n := pq.Len()
	pq.Swap(0, n-1)
	*pq = (*pq)[0 : n-1]

	cur := 0
	bh.siftDown(cur)
	// cidx1, cidx2 := getChildrenIdx(cur, pq.Len())

	// for cidx1 > 0 {
	// 	if cidx2 < 0 {
	// 		if !pq.Less(cur, cidx1) {
	// 			pq.Swap(cur, cidx1)
	// 			cur = cidx1
	// 			cidx1, cidx2 = getChildrenIdx(cur, pq.Len())
	// 		} else {
	// 			break
	// 		}
	// 	} else {
	// 		if pq.Less(cidx1, cidx2) {
	// 			if !pq.Less(cur, cidx1) {
	// 				pq.Swap(cur, cidx1)
	// 				cur = cidx1
	// 				cidx1, cidx2 = getChildrenIdx(cur, pq.Len())
	// 			} else {
	// 				break
	// 			}
	// 		} else {
	// 			if !pq.Less(cur, cidx2) {
	// 				pq.Swap(cur, cidx2)
	// 				cur = cidx2
	// 				cidx1, cidx2 = getChildrenIdx(cur, pq.Len())
	// 			} else {
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	return item

}

type Item struct {
	key   int
	count int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// func (pq *PriorityQueue) Push(x interface{}) {
// 	item := x.(*Item)
// 	*pq = append(*pq, item)
// }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

// func topKFrequent(nums []int, k int) []int {
// 	m := make(map[int]int)
// 	for _, n := range nums {
// 		m[n]++
// 	}
// 	q := PriorityQueue{}
// 	for key, count := range m {
// 		heap.Push(&q, &Item{key: key, count: count})
// 	}
// 	var result []int
// 	for len(result) < k {
// 		item := heap.Pop(&q).(*Item)
// 		result = append(result, item.key)
// 	}
// 	return result
// }

// using owne version of heap
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	h := new(BiHeap)

	for key, count := range m {
		h.Push(Item{key: key, count: count})
	}
	var result []int
	for len(result) < k {
		item := h.Pop()
		result = append(result, item.key)
	}
	return result
}

func Test() {
	nums := []int{1, 2, 2, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 5, 5, 8, 5, 8, 8, 8, 8, 8}
	// nums := []int{1, 1, 1, 2, 2, 3}
	result := topKFrequent(nums, 4)
	fmt.Println(result)
}
