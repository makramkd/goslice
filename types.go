package goslice

type Pair[T any, U any] struct {
	First  T
	Second U
}

func MakePair[T any, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{
		First:  first,
		Second: second,
	}
}
