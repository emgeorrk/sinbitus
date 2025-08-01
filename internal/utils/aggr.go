package utils

func FirstNonZero[T comparable](vals ...T) T {
	defaultVal := *new(T)
	for _, val := range vals {
		if val != defaultVal {
			return val
		}
	}
	return defaultVal
}

func Coalesce[T any](vals ...*T) *T {
	for _, val := range vals {
		if val != nil {
			return val
		}
	}
	return nil
}
