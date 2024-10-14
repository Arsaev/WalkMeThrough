package utils

// Find - generic find function that takes a slice and a function that returns a boolean value for each element
// in the slice and returns the first element for which the function returns true and the index of that element
// in the slice.
func Find[T any](slice []T, f func(T) bool) (T, int) {
	for i, v := range slice {
		if f(v) {
			return v, i
		}
	}
	var zero T
	return zero, -1
}
