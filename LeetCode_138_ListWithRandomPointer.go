package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//func copyRandomList(head *Node) *Node {
//	p :=head
//	var data []*Node
//	// 切片存每个节点
//	for p!=nil{
//		// 复制新的元素
//		newL :=new(Node)
//		newL = p
//		data = append(data,newL)
//		p=p.Next
//	}
//	// q 为新链表表 头
//	q :=data[0]
//	// 复制random索引
//	for i :=0;i<len(data)-1;i++{
//		// 顺序复制勾链
//		q.Next=data[i+1]
//		// Random 勾链
//		if q.Random!=nil{
//			q.Random =
//		}else{
//			q.Random =nil
//		}
//		q=q.Next
//	}
//	return nil
//}
func copyList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 增加新的节点
	for node := head; node != nil; node = node.Next.Next {
		// 在每个节点之后复制连接自己的复制节点
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	// 连接 Random
	for node := head; node != nil; node = node.Next.Next {
		// node.Next 为复制节点 A'
		if node.Random != nil {
			// node.Random.Next 表示连接到 Random 的复制节点
			node.Next.Random = node.Random.Next
		}
	}
	// 复制链表勾链
	NewHead := head.Next
	// node 表示指向链表
	for node := head; node != nil; node = node.Next {
		// nodeNew 指向复制的节点
		nodeNew := node.Next
		// 恢复原链表
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			// 指向下一个复制的节点
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return NewHead
}
