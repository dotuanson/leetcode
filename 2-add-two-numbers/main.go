package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var iterateL1Node, iterateL2Node, resultNode = l1, l2, &ListNode{}
	var headNode = resultNode
	var sumNode = 0
	for {
		if (iterateL1Node == nil) && (iterateL2Node == nil) {
			resultNode.Val = sumNode
			resultNode.Next = nil
			break
		} else {
			resultNode.Next = &ListNode{}
		}
		if iterateL1Node != nil {
			sumNode += iterateL1Node.Val
			iterateL1Node = iterateL1Node.Next
		}
		if iterateL2Node != nil {
			sumNode += iterateL2Node.Val
			iterateL2Node = iterateL2Node.Next
		}
		resultNode.Val = sumNode % 10
		sumNode /= 10
		resultNode = resultNode.Next
	}

	node := headNode
	for {
		if node.Next.Next == nil {
			break
		}
		node = node.Next
	}
	if node.Next.Val == 0 {
		node.Next = nil
	}

	//for node := headNode; node != nil; node = node.Next {
	//	log.Println(node.Val, node.Next)
	//}
	//
	//log.Println()
	return headNode
}
