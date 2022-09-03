package slices

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveByIndexStr(t *testing.T) {
	type testData struct {
		arr     []string
		element int
		want    error
	}

	tests := []testData{
		{[]string{"null", "one", "two", "four"}, 1, nil},
		{[]string{"0", "1", "2"}, 4, errors.New(fmt.Sprintf("index not found %v", 4))},
	}

	for _, test := range tests {
		got := RemoveByIndex(test.arr, test.element)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("RemoveByIndex([]T, int) error, want: %v, got: %v", test.want, got)
		}
	}
}

func TestRemoveByIndexInt(t *testing.T) {
	type testData struct {
		arr     []int
		element int
		want    error
	}

	tests := []testData{
		{[]int{0, 1, 2, 3, 4}, 0, nil},
		{[]int{0, 1, 2, 3, 4}, 10, errors.New(fmt.Sprint("index not found 10"))},
	}

	for _, test := range tests {
		got := RemoveByIndex(test.arr, test.element)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("RemoveByIndex([]T, int) error, want: %v, got: %v", test.want, got)
		}
	}
}
