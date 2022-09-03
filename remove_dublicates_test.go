package slices

import (
	"testing"
	"reflect"

)

func TestRemoveDublicatesStr(t *testing.T) {
  type testData struct {
    arr []string
    want []string
  }

  tests := []testData{
    {[]string{"one", "two", "three", "one"}, []string{"one", "two", "three"}},
    {[]string{"1", "2", "3", "4", "1", "2"}, []string{"1", "2", "3", "4"}},
    {[]string{}, []string{}},
  }

  for _, test := range tests {
    got := RemoveDublicatesStr(test.arr)
    if !reflect.DeepEqual(got, test.want) {
      t.Errorf("RemoveDublicatesStr([]string) []string, want: %v, got: %v\n", test.want, got)
    }
  }
}
