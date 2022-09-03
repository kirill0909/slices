package slices

import (
	"errors"
	"reflect"
	"testing"
	"fmt"
)

func TestRemoveByIndexStr(t *testing.T) {
  type testData struct {
    arr []string
    element int
    want error
  }

  tests := []testData{
    {[]string{"null", "one", "two", "four"}, 1, nil},
    {[]string{"0", "1", "2"}, 4, errors.New(fmt.Sprintf("slice does not contain an element with an index %v", 4))},
  }

  for _, test := range tests {
    got := RemoveByIndex(test.arr, test.element)
    if !reflect.DeepEqual(got, test.want) {
      t.Errorf("RemoveByIndex([]T, int) error, want: %v, got: %v", test.want, got)
    }
  }
}
