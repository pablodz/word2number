package utils

func IsItemInSlice(item string, slice []string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func InsertValueByIndexInSlice(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
