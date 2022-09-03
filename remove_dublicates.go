package slices

/*
RemoveDublicatesStr removes duplicatesw from []string
*/
func RemoveDublicatesStr(arr []string) []string {
  allKeys := make(map[string]bool)
  list := []string{}
  for _, item := range arr {
    if _, value := allKeys[item]; !value {
      allKeys[item] = true
      list = append(list, item)
    }
  }
  return list
}
