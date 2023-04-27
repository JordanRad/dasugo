package dasugo

type ReduceResult struct {
	Value float64
}

// Reduce applies a function against an accumulator and each element in the slice to reduce it to a single value.
//
// Parameters:
// arr: a slice of type T
// fn: a function that takes two parameters - an accumulator of type ReduceResult and a current value of type T, and returns a new accumulator of type ReduceResult
// initialValue: an initial accumulator value of type ReduceResult
//
// Return:
// a ReduceResult struct that represents the final accumulator value
//
// Example:
//
//	products := []Product{{"apple", 1.0}, {"banana", 2.0}, {"orange", 3.0}}
//	totalPrice := Reduce(products, func(acc ReduceResult, cv Product) ReduceResult {
//	  acc.Value += cv.Price
//	  return acc
//	}, ReduceResult{0})
//	fmt.Println(totalPrice.Value) // Output: 6.0
func Reduce[T any](arr []T, fn func(acc ReduceResult, cv T) ReduceResult, initialValue ReduceResult) ReduceResult {
	result := initialValue
	for _, item := range arr {
		result = fn(result, item)
	}
	return result
}
