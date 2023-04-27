package dasugo_test

import (
	"github.com/JordanRad/dasugo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Product struct {
	Name  string
	Price float32
}

var _ = Describe("Dasugo", func() {
	Describe("Reduce", func() {
		It("should apply the mapping function sum the products prices", func() {
			input := []Product{{Name: "A", Price: 75.5}, {Name: "B", Price: 15.5}, {Name: "C", Price: 9.00}}
			expected := 100.00

			output := dasugo.Reduce(input, func(acc dasugo.ReduceResult, cv Product) dasugo.ReduceResult {
				return dasugo.ReduceResult{
					Value: float64(cv.Price) + acc.Value,
				}
			}, dasugo.ReduceResult{Value: 0})

			Expect(output.Value).To(Equal(expected))
		})
	})
})
