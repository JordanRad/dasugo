package dasugo_test

import (
	"github.com/JordanRad/dasugo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dasugo", func() {
	Describe("Contains", func() {
		It("should return true if the input array contains the wanted entry", func() {
			input := []int{1, 2, 3, 40, 50}
			output := dasugo.Contains(input, 50)

			Expect(output).To(BeTrue())
		})
		It("should return false if the input array does NOT contain the wanted entry", func() {
			input := []int{1, 2, 3, 40, 50}
			output := dasugo.Contains(input, 250)

			Expect(output).To(BeFalse())
		})
	})
})
