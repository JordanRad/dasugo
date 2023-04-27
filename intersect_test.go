package dasugo_test

import (
	"github.com/JordanRad/dasugo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dasugo", func() {
	Describe("Intersect", func() {
		It("should return the common entries", func() {
			input1 := []int{1, 2, 3, 40, 50}
			input2 := []int{1, 100, 2}
			expected := []int{1, 2}
			output := dasugo.Intersect(input1, input2)

			Expect(output).To(Equal(expected))
		})
		It("should return an empty array of strings if there are no common elements", func() {
			input1 := []int{1, 2, 3, 40, 50}
			input2 := []int{1897, 100, 7652}

			output := dasugo.Intersect(input1, input2)
			Expect(output).To(BeEmpty())
		})
	})
})
