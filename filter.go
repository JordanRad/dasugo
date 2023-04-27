package dasugo

// Filter returns a new slice containing all elements in `arr` that satisfy the condition specified by `fn`.
// The function `fn` takes an element of type `T` as input and returns a boolean value.
// If `fn(item)` returns `true` for an element `item` in `arr`, that element is included in the returned slice.
// If `fn(item)` returns `false`, that element is excluded from the returned slice.
// The returned slice has the same element type as `arr`.
func Filter[T any](arr []T, fn func(item T) bool) []T {
	result := make([]T, 0)
	for _, item := range arr {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}
