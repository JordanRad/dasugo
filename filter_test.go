package dasugo_test

import (
	"github.com/JordanRad/dasugo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dasugo", func() {
	Describe("Filter", func() {
		It("should apply the mapping function to exclude any number below 10", func() {
			input := []int{1, 2, 3, 40, 50}
			expected := []int{40, 50}

			output := dasugo.Filter(input, func(item int) bool {
				return item >= 10
			})

			Expect(output).To(Equal(expected))
		})
	})
})
