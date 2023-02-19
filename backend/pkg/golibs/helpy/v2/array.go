package helpy

import "golang.org/x/exp/constraints"

func ExistsInArray[T comparable](value T, slice []T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func AddPrefixInArrayString(list []string, prefix string) []string {
	for i := range list {
		list[i] = prefix + list[i]
	}
	return list
}

type SortSlice[T constraints.Ordered] []T

func (x SortSlice[T]) Len() int           { return len(x) }
func (x SortSlice[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x SortSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
