package main

import (
	"fmt"

	"github.com/benjgrad/COMP/src/app/datastructures/linkedlist"
)

func main() {

	list := linkedlist.LList{}

	list.Prepend(&linkedlist.Node{Value: 1})
	list.Prepend(&linkedlist.Node{Value: 2})
	list.Prepend(&linkedlist.Node{Value: 6})
	list.Prepend(&linkedlist.Node{Value: 9})

	//Print with formatting
	fmt.Printf("Length of the list is: %v\n", list.Size)

}
