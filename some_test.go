package dasugo_test

import (
	"github.com/JordanRad/dasugo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Student struct {
	Name string
	Age  int
}

var _ = Describe("Dasugo", func() {
	Describe("Some", func() {
		It("should return true if at least one entry meets a given condition", func() {
			input := []Student{{Name: "A", Age: 19}, {Name: "B", Age: 14}, {Name: "C", Age: 11}}
			output := dasugo.Some(input, func(s Student) bool {
				return s.Age >= 18
			})

			Expect(output).To(BeTrue())
		})

		It("should return false if no entry meets a given condition", func() {
			input := []Student{{Name: "A", Age: 19}, {Name: "B", Age: 14}, {Name: "C", Age: 11}}
			output := dasugo.Some(input, func(s Student) bool {
				return s.Age > 218
			})

			Expect(output).To(BeFalse())
		})
	})
})
