package slices

// RemoveDublicatesStr removes duplicatesw from []T
func RemoveDublicates[T comparable](arr []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range arr {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
