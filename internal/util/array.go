package util

func Includes[K comparable](s []K, v K) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}

	return false
}
