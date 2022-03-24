package goslice_test

import (
	"testing"

	"github.com/makramkd/goslice"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	type s struct {
		a int
		b int32
	}
	coll := []s{{a: 1, b: 2}, {a: 2, b: 3}, {a: 3, b: 4}}
	theAs := goslice.Map(coll, func(elem s) int {
		return elem.a
	})
	require.ElementsMatch(t, []int{1, 2, 3}, theAs)

	theBs := goslice.Map(coll, func(elem s) int32 {
		return elem.b
	})
	require.ElementsMatch(t, []int32{2, 3, 4}, theBs)
}

func TestFilter(t *testing.T) {
	type s struct {
		a int
		b int32
	}
	coll := []s{{a: 1, b: 2}, {a: 2, b: 3}, {a: 3, b: 4}}
	evenAs := goslice.Filter(coll, func(elem s) bool {
		return elem.a%2 == 0
	})
	require.ElementsMatch(t, []s{{a: 2, b: 3}}, evenAs)
}

func TestFilterNot(t *testing.T) {
	type s struct {
		a int
		b int32
	}
	coll := []s{{a: 1, b: 2}, {a: 2, b: 3}, {a: 3, b: 4}}
	oddAs := goslice.FilterNot(coll, func(elem s) bool {
		return elem.a%2 == 0
	})
	require.ElementsMatch(t, []s{{a: 1, b: 2}, {a: 3, b: 4}}, oddAs)
}

func TestReduce(t *testing.T) {
	// Simple case: accumulator and collection type are the same
	coll := []int{1, 2, 3, 4, 5}
	sum := goslice.Reduce(coll, func(a int, b int) int {
		return a + b
	}, 0)
	require.Equal(t, 1+2+3+4+5, sum)

	// general case: accumulator and collection type are different
	type s struct {
		a int
		b int32
	}
	coll2 := []s{{a: 1, b: 2}, {a: 2, b: 3}, {a: 3, b: 4}}
	abSum := goslice.Reduce(coll2, func(accum int64, elem s) int64 {
		return accum + int64(elem.a) + int64(elem.b)
	}, int64(0))
	require.Equal(t, int64(1+2+2+3+3+4), abSum)
}
