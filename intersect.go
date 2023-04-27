package dasugo

func Intersect[T comparable](arr1, arr2 []T) []T {
	elementsLookUp := arr1
	result := make([]T, 0)
	for _, item := range arr2 {
		if Contains(elementsLookUp, item) {
			result = append(result, item)
		}
	}
	return result
}
