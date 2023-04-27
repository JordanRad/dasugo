package dasugo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDasugo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dasugo Suite")
}
