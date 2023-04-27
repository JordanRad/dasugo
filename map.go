package dasugo

// Map applies the function fn to each element in the input slice arr
// and returns a new slice with the result of applying the function to
// each element. The type of the input slice arr and the type of its
// elements are represented by the type parameter T, which is constrained
// by the "any" keyword to allow any type.
//
// The function fn takes an element of type T as its argument and returns
// an element of the same type T after applying some operation to it.
//
// The returned slice has the same length as the input slice arr and its
// elements have the same type T as the input slice.
func Map[T any](arr []T, fn func(item T) T) []T {
	result := make([]T, len(arr))
	for i, item := range arr {
		result[i] = fn(item)
	}
	return result
}
