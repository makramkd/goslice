package goslice

import "golang.org/x/exp/constraints"

// SumSlice returns the sum of all the elements in the given slice.
// No attempt is made to detect overflow.
func SumSlice[T constraints.Ordered](elems []T) T {
	var sum T
	for _, elem := range elems {
		sum += elem
	}
	return sum
}

// ProdSlice returns the product of all the elemenst in the given slice.
// No attempt is made to detect overflow.
func ProdSlice[T constraints.Integer | constraints.Float](elems []T) T {
	var prod T = 1
	for _, elem := range elems {
		prod *= elem
	}
	return prod
}

// AllOf returns true if and only if all of the elements in the slice satisfy the given predicate.
func AllOf[T any](elems []T, pred func(elem T) bool) bool {
	var all bool = true
	for _, elem := range elems {
		all = all && pred(elem)
	}
	return all
}

// AnyOf returns true if and only if at least one of the elements in the slice satisfy the given predicate.
func AnyOf[T any](elems []T, pred func(elem T) bool) bool {
	var anyOf bool = false
	for _, elem := range elems {
		anyOf = anyOf || pred(elem)
	}
	return anyOf
}

// NoneOf returns true if and only if none of the elements in the slice satisfy the given predicate.
func NoneOf[T any](elems []T, pred func(elem T) bool) bool {
	return !AnyOf(elems, pred)
}

// CountIf returns the number of elements in the slice that satisfy the given predicate.
func CountIf[T any](elems []T, pred func(elem T) bool) (count int) {
	for _, elem := range elems {
		if pred(elem) {
			count++
		}
	}
	return
}

// FindIf returns the index to the first element in the slice that satisfies the given predicate.
// If no such element exists, -1 is returned.
func FindIf[T any](elems []T, pred func(elem T) bool) int {
	for i, elem := range elems {
		if pred(elem) {
			return i
		}
	}
	return -1
}

// Merge merges the two provided sorted slices into a single sorted slice.
func Merge[T constraints.Ordered](left, right []T) (sorted []T) {
	sorted = make([]T, len(left)+len(right))

	i := 0
	j := 0
	k := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted[k] = left[i]
			i++
			k++
		} else {
			sorted[k] = right[j]
			j++
			k++
		}
	}

	for i < len(left) {
		sorted[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		sorted[k] = right[j]
		k++
		j++
	}

	return
}

// BinarySearch returns the index of the given target in the provided sorted elems slice.
// If the target is not in the sorted slice, -1 is returned.
func BinarySearch[T constraints.Ordered](elems []T, target T) (index int) {
	lower, upper := 0, len(elems)-1
	for lower <= upper {
		mid := lower + (upper-1)/2

		if elems[mid] == target {
			return mid
		} else if elems[mid] < target {
			lower = mid + 1
		} else {
			upper = mid - 1
		}
	}

	return -1
}
