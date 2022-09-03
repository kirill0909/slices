package slices

import (
	"reflect"
	"testing"
)

func TestRemoveDublicatesStr(t *testing.T) {
	type testData struct {
		arr  []string
		want []string
	}

	tests := []testData{
		{[]string{"one", "two", "three", "one"}, []string{"one", "two", "three"}},
		{[]string{"1", "2", "3", "4", "1", "2"}, []string{"1", "2", "3", "4"}},
		{[]string{}, []string{}},
	}

	for _, test := range tests {
		got := RemoveDublicates(test.arr)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("RemoveDublicatesStr([]string) []string, want: %v, got: %v\n", test.want, got)
		}
	}
}

func TestRemoveDublicatesInt(t *testing.T) {
	type testData struct {
		arr  []int
		want []int
	}

	tests := []testData{
		{[]int{1, 2, 3, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 1, 2}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		got := RemoveDublicates(test.arr)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("RemoveDublicatesStr([]string) []string, want: %v, got: %v\n", test.want, got)
		}
	}
}
