package linkedlist
import "fmt"
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

func (list *LList) Print() {
	if (list.Size > 0){
		currentNode := *list.Head
		fmt.Println("\nList Size Value:", list.Size)
		for i := 0; i < list.Size; i++ {
			fmt.Println("Node Value:", currentNode.Value, "Next:", currentNode.Next)
			if currentNode.Next != nil {
				currentNode = *currentNode.Next
			}
		}
		fmt.Println("")
	}
}