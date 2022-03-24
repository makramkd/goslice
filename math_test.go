package goslice_test

import (
	"testing"

	"github.com/makramkd/goslice"
	"github.com/stretchr/testify/require"
)

func TestMax(t *testing.T) {
	x := int(1)
	y := int(2)
	expected := int(2)
	actual := goslice.Max(x, y)
	require.Equal(t, expected, actual)
}

func TestMin(t *testing.T) {
	x := int64(1)
	y := int64(2)
	expected := int64(1)
	actual := goslice.Min(x, y)
	require.Equal(t, expected, actual)
}
