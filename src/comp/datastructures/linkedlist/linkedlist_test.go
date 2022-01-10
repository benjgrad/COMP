package linkedlist

import (
	"github.com/benjgrad/COMP/src/comp/datastructures/linkedlist"
	"testing"
)

func TestPrepend(t *testing.T) {
	var tests = []struct {
		input  int
		output int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
	}

	for _, test := range tests {
		sut := linkedlist.LList{}
		for i := 0; i < test.input; i++ {
			sut.Prepend(&linkedlist.Node{Value: test.input - i})
		}

		if sut.Size != test.output {
			t.Error("Test failed: {} expected, received {}", test.output, sut.Size)
		}

		currentNode := *sut.Head
		t.Log("\nList Size Value:", sut.Size)
		for i := 0; i < sut.Size; i++ {
			t.Log("Node Value:", currentNode.Value, "Next:", currentNode.Next)
			if currentNode.Value != i+1 {
				t.Error("Test failed: {} expected, received {}", test.output, sut.Size)
			}
			if currentNode.Next != nil {
				currentNode = *currentNode.Next
			}
		}
	}
}
