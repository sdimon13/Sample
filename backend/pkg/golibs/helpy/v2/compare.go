package helpy

func Compare[T comparable](a, b *T) bool {
	switch {
	case a == nil && b == nil:
	case a != nil && b != nil:
		if *a != *b {
			return false
		}
	default:
		return false
	}
	return true
}
func CompareSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, vala := range a {
		found := false
		for _, valb := range b {
			if vala == valb {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}
