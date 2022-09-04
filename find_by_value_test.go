package slices

import (
	"testing"
)

func TestFindByValueStr(t *testing.T) {
	type testData struct {
		arr   []string
		value string
		want  bool
	}

	tests := []testData{
		{[]string{"null", "one", "two", "three"}, "one", true},
		{[]string{"null", "one", "two", "three"}, "1", false},
	}

	for _, test := range tests {
		_, got := FindByValue(test.arr, test.value)
		if got != test.want {
			t.Errorf("FindByValue([]T, T)(int, bool) want: %v, got: %v", test.want, got)
		}
	}
}

func TestFindByValuerInt(t *testing.T) {
	type testData struct {
		arr   []int
		value int
		want  bool
	}

	tests := []testData{
		{[]int{0, 1, 2, 3}, 3, true},
		{[]int{0, 1, 2}, 10, false},
	}

	for _, test := range tests {
		_, got := FindByValue(test.arr, test.value)
		if got != test.want {
			t.Errorf("FindByValue([]T, T)(int, bool) want: %v, got: %v", test.want, got)
		}
	}
}
