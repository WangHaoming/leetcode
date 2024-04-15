package reversekgroup

type ListNode struct {
	val  int
	next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	
	return head
}

func makeList(nums []int) *ListNode {
	var head *ListNode
	var lastNode *ListNode
	for i := 0; i < len(nums); i++ {
		node := new(ListNode)
		node.val = nums[i]
		node.next = nil
		if head == nil {
			head = node
			lastNode = node
		} else {
			lastNode.next = node
			lastNode = node
		}
	}
	return head
}

func printList(head *ListNode) {
	lastNode := head
	for {
		print(lastNode.val, ",")
		if lastNode.next != nil {
			lastNode = lastNode.next
		} else {
			break
		}
	}
	println()
}

func Test() {
	nums := []int{1, 2, 3, 4, 5}
	head := makeList(nums)

	reversedhead := reverseKGroup(head, 2)
	printList(reversedhead)

}
