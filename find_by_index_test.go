package slices

import (
  "testing"
)

func TestFindByIdStr(t *testing.T) {
  type testData struct {
    arr []string
    index int
    want bool
  }

  tests := []testData{
    {[]string{"null", "one", "two", "three"}, 1, true},
    {[]string{"0", "1", "2", "3"}, 2, true},
    {[]string{"0", "1", "2", "three", "four"}, 10, false},
  }

  for _, test := range tests {
    got := FindByIndex(test.arr, test.index)
    if got != test.want {
      t.Errorf("FindByIndex[T]([]T, int)(int, bool) want: %v, got: %v", test.want, got)
    }
  }
}

func TestFindByIdInt(t *testing.T) {
  type testData struct {
    arr []int
    index int
    want bool
  }

  tests := []testData{
    {[]int{0, 1, 2, 3}, 1, true},
    {[]int{}, 0, false},
    {[]int{10, 20, 30}, 4, false},
  }

  for _, test := range tests {
    got := FindByIndex(test.arr, test.index)
    if got != test.want {
      t.Errorf("FindByIndex[T]([]T, int)(int, bool) want: %v, got: %v", test.want, got)
    }
  }
}
