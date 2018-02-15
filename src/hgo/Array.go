package hgo

// Example: [3 4 0 1 2], 2
func GetShiftedIndexes(length, index int) (array []int) {
	array = make([]int, length)
	for i := range array {
		if i == index {
			array[i] = 0
		} else if index < i {
			array[i] = i - index
		} else {
			array[i] = len(array) + i - index
		}
	}
	return
}
