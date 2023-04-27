package dasugo

// Some returns true if at least one element in the slice satisfies the given predicate function, false otherwise.
// The predicate function should take an element of type T as input and return a boolean value.
// The function iterates through the slice and calls the predicate function on each element until a true value is returned.
// If a true value is found, the function immediately returns true, otherwise it returns false.
// Example:
//
//	arr := []int{1, 2, 3}
//	result := Some(arr, func(item int) bool { return item > 2 })
//	result == true
func Some[T any](arr []T, fn func(item T) bool) bool {
	for _, item := range arr {
		if fn(item) {
			return true
		}
	}
	return false
}
