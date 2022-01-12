package algorithms

func BinarySearch (arr []int, target int) int {
	var pivot, left, right int = 0, 0, len(arr) - 1
	for ; left<=right ; {
		pivot = (right + left) / 2
		if target == arr[pivot] {
			return pivot
		} else if target > arr[pivot] {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return -1
}