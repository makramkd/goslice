package goslice

// Map applies the given function to each slice element and returns the resultant slice.
func Map[T any, U any](elems []T, f func(T) U) (r []U) {
	r = make([]U, len(elems))
	for i := range elems {
		r[i] = f(elems[i])
	}
	return
}

// Filter returns all elements that satisfy the given predicate in a new slice.
func Filter[T any](elems []T, p func(T) bool) (r []T) {
	for _, e := range elems {
		if p(e) {
			r = append(r, e)
		}
	}
	return
}

// Filter returns all elements that do NOT satisfy the given predicate in a new slice.
func FilterNot[T any](elems []T, p func(T) bool) (r []T) {
	for _, e := range elems {
		if !p(e) {
			r = append(r, e)
		}
	}
	return
}

// Reduce reduces the given slice to a single value using the given function and
// initial value.
func Reduce[T any, U any](elems []T, f func(a U, b T) U, initial U) (r U) {
	r = initial
	for _, e := range elems {
		r = f(r, e)
	}
	return
}
