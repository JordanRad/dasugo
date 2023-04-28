package dasugo_test

import (
	"github.com/JordanRad/dasugo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dasugo", func() {
	Describe("Find", func() {
		It("should return a pointer to entry if it exists", func() {
			input := []Product{{Name: "A", Price: 10.00}, {Name: "B", Price: 40.00}}
			expected := Product{Name: "A", Price: 10.00}

			output := dasugo.Find(input, func(t Product) bool {
				return t.Price == 10.00
			})

			Expect(output).To(Not(BeNil()))
			Expect(output.Name).To(Equal(expected.Name))
			Expect(output.Price).To(Equal(expected.Price))
		})

		It("should return nil if the searched entry cannot be found", func() {
			input := []Product{{Name: "A", Price: 10.00}, {Name: "B", Price: 40.00}}

			output := dasugo.Find(input, func(t Product) bool {
				return t.Price >= 190.00
			})

			Expect(output).To(BeNil())
		})
	})

	Describe("FindIndex", func() {
		It("should return an integer with the index of the found entry if it exists", func() {
			input := []Product{{Name: "A", Price: 10.00}, {Name: "B", Price: 40.00}}
			expected := 0

			output := dasugo.FindIndex(input, func(t Product) bool {
				return t.Price == 10.00
			})

			Expect(output).To(Equal(expected))
		})

		It("should return -1 if the searched entry cannot be found", func() {
			input := []Product{{Name: "A", Price: 10.00}, {Name: "B", Price: 40.00}}

			output := dasugo.FindIndex(input, func(t Product) bool {
				return t.Price >= 190.00
			})

			Expect(output).To(Equal(-1))
		})
	})
})
