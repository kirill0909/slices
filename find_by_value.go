package slices

// FindByValue search element in slice []T by value
func FindByValue[T comparable](arr []T, value T) (int, bool) {
  for i, v := range arr {
    if v == value {
      return i, true
    }
  }

  return -1, false
}
