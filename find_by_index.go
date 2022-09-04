package slices

// FindByIndex search element in slice []T by index
func FindByIndex[T comparable](arr []T, index int)  bool{
  for i := range arr {
    if i == index {
      return true
    }
  }
  return false
}
