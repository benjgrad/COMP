package algorithms

import (
	"testing"
	"github.com/benjgrad/COMP/src/comp/algorithms"
)

func TestBinarySearch_Int(t *testing.T) {
	array := []int{-1,0,3,5,9,12}
	target := 9
	output := algorithms.BinarySearch(array, target)
	if output != 4 {
		t.Error("Test failed: {} expected, received {}", output, 4)
	}
}