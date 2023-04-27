package dasugo_test

import (
	"github.com/JordanRad/dasugo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dasugo", func() {

	Describe("Map", func() {
		It("should apply the mapping function to multiply by 10 to all elements in the slice", func() {
			input := []int{1, 2, 3}
			expected := []int{10, 20, 30}

			output := dasugo.Map(input, func(item int) int {
				return item * 10
			})

			Expect(output).To(Equal(expected))
		})
	})
})
