package dasugo

// Contains checks if an element of type T is contained within a slice of type []T. It returns a boolean value indicating whether or not the searched item is present in the given slice.

// arr []T: a slice of elements of type T
// searchedItem T: the element of type T to search for in the slice
// bool: returns true if the searched item is found in the slice, false otherwise
// The type parameter T must satisfy the comparable constraint, which means that the == operator must be defined for values of type T.
func Contains[T comparable](arr []T, searchedItem T) bool {
	for _, item := range arr {
		if item == searchedItem {
			return true
		}
	}
	return false
}
