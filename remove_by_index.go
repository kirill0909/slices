package slices

import (
	"errors"
	"fmt"
)

// RemoveByIndex remove element from slice []T by index
func RemoveByIndex[T comparable](arr *[]T, index int) error {
	if len(*arr) >= index {
		(*arr) = append((*arr)[:index], (*arr)[index+1:]...)
		return nil
	}

	return errors.New(fmt.Sprintf("index not found %v", index))
}
