package slices

import (
	"errors"
	"fmt"
)

// RemoveByElement removes element from []T by value
func RemoveByValue[T comparable](arr []T, value T) error {
	for i, v := range arr {
	  if v == value {
	    arr = append(arr[:i], arr[i+1:]...)
	    return nil
	  }
	}
	return errors.New(fmt.Sprintf("value not found %v", value))
}
