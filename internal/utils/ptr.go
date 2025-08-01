package utils

func GetPtr[T any](val T) *T {
	return &val
}

func SafeDeref[T any](val *T) T {
	if val != nil {
		return *val
	}

	return *new(T)
}
