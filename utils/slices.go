package utils

// Return value of last element.
func LastElement[T any](slices []T) T {
	return slices[LastIndex(slices)]
}

// Return last index of slices
func LastIndex[T any](slices []T) int {
	return len(slices) - 1
}
