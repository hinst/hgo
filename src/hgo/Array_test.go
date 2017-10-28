package hgo

import "testing"

func TestSetShiftedIndexes(t *testing.T) {
	{
		var target = []int{3, 4, 0, 1, 2}
		var source = GetShiftedIndexes(5, 2)
		Assert(len(source) == len(target))
		for i := range source {
			Assert(source[i] == target[i])
		}
	}
	{
		var target = []int{2, 0, 1}
		var source = GetShiftedIndexes(3, 2)
		Assert(len(source) == len(target))
		for i := range source {
			println(source[i])
		}
		for i := range source {
			Assert(source[i] == target[i])
		}
	}
}
