package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got '%d' expected '%d' given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// Use new slices.Equal instead of reflect.DeepEqual
	if !slices.Equal(got, want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	// Use new slices.Equal instead of reflect.DeepEqual
	if !slices.Equal(got, want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
