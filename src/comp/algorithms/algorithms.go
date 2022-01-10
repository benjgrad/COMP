package algorithms

func BinarySearch (arr []int, target int) int {
	var pivot, left, right int = 0, 0, len(arr) - 1
	for ; left<=right ; {
		pivot = left + (right - left) / 2
		if target == arr[pivot] {
			return pivot
		} else if target > pivot {
			left = pivot
		} else {
			right = pivot
		}
	}
	return -1
}