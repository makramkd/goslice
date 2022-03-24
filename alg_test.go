package goslice_test

import (
	"testing"

	"github.com/makramkd/goslice"
	"github.com/stretchr/testify/require"
)

func TestSumSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	expected1 := 1 + 2 + 3 + 4 + 5
	actual1 := goslice.SumSlice(s1)
	require.Equal(t, expected1, actual1)

	s2 := []uint64{2, 3, 4, 5, 6}
	expected2 := uint64(2 + 3 + 4 + 5 + 6)
	actual2 := goslice.SumSlice(s2)
	require.Equal(t, expected2, actual2)
}

func TestProdSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	expected1 := 1 * 2 * 3 * 4 * 5
	actual1 := goslice.ProdSlice(s1)
	require.Equal(t, expected1, actual1)

	s2 := []uint64{2, 3, 4, 5, 6}
	expected2 := uint64(2 * 3 * 4 * 5 * 6)
	actual2 := goslice.ProdSlice(s2)
	require.Equal(t, expected2, actual2)
}

func TestAllOf(t *testing.T) {
	isEven := func(i int) bool { return i%2 == 0 }
	require.False(t, goslice.AllOf([]int{1, 2, 3, 4, 5}, isEven))
	require.True(t, goslice.AllOf([]int{2, 4, 6, 8, 10}, isEven))

	isOdd := func(i int64) bool { return i%2 == 1 }
	require.False(t, goslice.AllOf([]int64{1, 2, 3, 4, 5}, isOdd))
	require.True(t, goslice.AllOf([]int64{1, 3, 5, 7, 9}, isOdd))
}

func TestAnyOf(t *testing.T) {
	isEven := func(i int) bool { return i%2 == 0 }
	require.True(t, goslice.AnyOf([]int{1, 2, 3, 4, 5}, isEven))  // some even
	require.False(t, goslice.AnyOf([]int{1, 3, 5, 7, 9}, isEven)) // all odd

	isOdd := func(i int64) bool { return i%2 == 1 }
	require.True(t, goslice.AnyOf([]int64{1, 2, 3, 4, 5}, isOdd))   // some odd
	require.False(t, goslice.AnyOf([]int64{2, 4, 6, 8, 10}, isOdd)) // all even
}

func TestNoneOf(t *testing.T) {
	isEven := func(i int) bool { return i%2 == 0 }
	require.True(t, goslice.NoneOf([]int{1, 3, 5, 7, 9, 11}, isEven)) // none even
	require.False(t, goslice.NoneOf([]int{1, 2, 5, 7, 9}, isEven))    // some even

	isOdd := func(i int64) bool { return i%2 == 1 }
	require.True(t, goslice.NoneOf([]int64{2, 4, 6, 8, 10}, isOdd))  // none odd
	require.False(t, goslice.NoneOf([]int64{2, 3, 6, 8, 10}, isOdd)) // some even
}

func TestCountIf(t *testing.T) {
	isEven := func(i int) bool { return i%2 == 0 }
	require.Equal(t, 3, goslice.CountIf([]int{2, 4, 6, 9, 9}, isEven))
	isOdd := func(i int64) bool { return i%2 == 1 }
	require.Equal(t, 2, goslice.CountIf([]int64{1, 1, 2, 2, 2, 2}, isOdd))
}

func TestFindIf(t *testing.T) {

}

func TestMerge(t *testing.T) {
	left := []int{1, 4, 8, 9, 13}
	right := []int{0, 3, 4, 5, 12, 15, 18}
	expected := []int{0, 1, 3, 4, 4, 5, 8, 9, 12, 13, 15, 18}
	actual := goslice.Merge(left, right)
	require.ElementsMatch(t, expected, actual)
}

func TestBinarySearch(t *testing.T) {
	arr := []int{0, 1, 3, 4, 4, 5, 8, 9, 12, 13, 15, 18}
	target := 4
	expected := 3
	actual := goslice.BinarySearch(arr, target)
	require.Equal(t, expected, actual)

	target = 2
	expected = -1
	actual = goslice.BinarySearch(arr, target)
	require.Equal(t, expected, actual)
}
