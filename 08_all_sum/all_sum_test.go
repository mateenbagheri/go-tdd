package main

import (
	"reflect"
	"testing"
)

func TestAllSum(t *testing.T) {
	t.Run("testing with no array input", func(t *testing.T) {
		wanted := []int{}
		got := allSum()
		assertCorrectMessageAllSum(t, got, wanted)
	})

	t.Run("testing with 1 array input", func(t *testing.T) {
		wanted := []int{10}
		got := allSum([]int{3, 7})
		assertCorrectMessageAllSum(t, got, wanted)
	})

	t.Run("testing with multiple array inputs", func(t *testing.T) {
		wanted := []int{11, 9}
		got := allSum([]int{2, 9}, []int{9})
		assertCorrectMessageAllSum(t, got, wanted)
	})

	t.Run("see if empty arrays are handled", func(t *testing.T) {
		wanted := []int{0, 9}
		got := allSum([]int{}, []int{3, 3, 3})
		assertCorrectMessageAllSum(t, got, wanted)
	})
}

func assertCorrectMessageAllSum(t testing.TB, got []int, wanted []int) {
	t.Helper()
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("wanted %v but instead, got %v", wanted, got)
	}
}
