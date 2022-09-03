package slices

import (
	"reflect"
	"testing"
	"errors"
	"fmt"
)

func TestRemoveByValueStr(t *testing.T) {
  type testData struct {
    arr[] string
    value string
    want error
  }

  tests := []testData{
    {[]string{"null", "one", "two", "three"}, "null", nil},
    {[]string{"null", "one", "two", "three"}, "eight", errors.New(fmt.Sprint("value not found eight"))},
  }

  for _, test := range tests {
    got := RemoveByValue(test.arr, test.value)
    if !reflect.DeepEqual(got, test.want) {
      t.Errorf("RemoveByValue([]T, T) error, want: %v, got: %v", test.want, got)
    }
  }
}

func TestRemoveByValueInt(t *testing.T) {
  type testData struct {
    arr[]int 
    value int
    want error
  }

  tests := []testData{
    {[]int{0, 1, 2, 3}, 0, nil},
    {[]int{0,1,2,3,4,5}, 8, errors.New(fmt.Sprint("value not found 8"))},
  }

  for _, test := range tests {
    got := RemoveByValue(test.arr, test.value)
    if !reflect.DeepEqual(got, test.want) {
      t.Errorf("RemoveByValue([]T, T) error, want: %v, got: %v", test.want, got)
    }
  }
}
