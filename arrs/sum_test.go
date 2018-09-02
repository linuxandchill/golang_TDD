package arrs

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("variable length slice", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		got := Sum(nums)
		expected := 10

		if got != expected {
			t.Errorf("got %d, expected %d give %v", got, expected, nums)
		}
	})

}

func TestSumAll(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{3, 2}

	got := SumAll(arr1, arr2)
	want := []int{6, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v, given %v and %v", got, want, arr1, arr2)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("sum tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}
		checkSums(t, got, want)
	})

	t.Run("run even when empty slices passed in", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4})
		want := []int{0, 4}
		checkSums(t, got, want)
	})
}
