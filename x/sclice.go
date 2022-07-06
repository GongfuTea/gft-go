package x

func Contains[T comparable](s []T, it T) bool {

	for _, v := range s {
		if v == it {
			return true
		}
	}
	return false
}
