package dasugo

// Intersect returns a new slice that contains only the elements that exist in both arr1 and arr2. The function takes two slices of type T and returns a new slice of type T.
//
// The T type must be comparable, which means that it must be possible to determine equality between two values of the type using the == operator.
//
// The function first creates a lookup slice (elementsLookUp) from arr1. Then, it iterates over arr2 and appends the items that exist in the lookup slice to the result slice.
func Intersect[T comparable](arr1, arr2 []T) []T {
	elementsLookUp := make(map[T]struct{}, len(arr1))
	for _, item := range arr1 {
		elementsLookUp[item] = struct{}{}
	}

	result := make([]T, 0, len(arr2))
	for _, item := range arr2 {
		_, found := elementsLookUp[item]
		if found {
			result = append(result, item)
		}
	}

	return result
}
