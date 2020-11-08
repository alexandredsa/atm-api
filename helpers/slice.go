package helpers

import "sort"

//Slice has utils functions to
type Slice struct {
}

//SortDesc sort slice of values int8 descending
func (s Slice) SortDesc(values []int16) {
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
}
