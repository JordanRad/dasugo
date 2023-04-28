package dasugo

// Find function takes a slice of type T and a function that takes an argument of type T and returns a bool. It returns a pointer to the first element of the slice for which the function returns true, or nil if no such element is found.
//
// arr []T: The slice of type T to search through.
//
// fn func(T) bool: The function to use to test each element of the slice. The function takes an argument of type T and returns a bool.
//
// *T: A pointer to the first element of the slice for which the function returns true, or nil if no such element is found.
func Find[T any](arr []T, fn func(T) bool) *T {
	for _, v := range arr {
		if fn(v) {
			return &v
		}
	}
	return nil
}

// FindIndex function takes a slice of type T and a function that takes an argument of type T and returns an int. It returns the index to the first element of the slice for which the function returns true, or -1 if no such element is found.
//
// arr []T: The slice of type T to search through.
//
// fn func(T) bool: The function to use to test each element of the slice. The function takes an argument of type T and returns a bool.
//
// int: Index to the first element of the slice for which the function returns true, or -1 if no such element is found.
func FindIndex[T any](arr []T, fn func(T) bool) int {
	for idx, v := range arr {
		if fn(v) {
			return idx
		}
	}
	return -1
}
