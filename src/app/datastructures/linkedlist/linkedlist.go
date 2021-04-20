package linkedlist

type Node struct {
	Value int
	Next  *Node
}

type LList struct {
	Size int
	Head *Node
}

func (list *LList) Prepend(node *Node) LList {

	node.Next = list.Head
	list.Head = node
	list.Size++
	return *list
}
