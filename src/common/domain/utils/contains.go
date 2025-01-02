package utils

func Contains[T comparable](slice []T, item T) bool {
	res := false
	for _, s := range slice {
		if s == item {
			res = true
			break
		}
	}
	return res
}
