package helpy

func Allocate[T any](val T) *T {
	return &val
}

func Unref[T any](ptr *T) (val T) {
	if ptr != nil {
		return *ptr
	}
	return
}
