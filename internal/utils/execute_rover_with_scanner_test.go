package utils

import (
	"bufio"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils: ExecuteRoverWithScanner", func() {
	When("first line != map size", func() {
		It("should return error invalid map size", func() {
			scanner := bufio.NewScanner(strings.NewReader("hello"))

			err := ExecuteRoverWithScanner(scanner, func(_ string) {})
			Expect(err.Error()).To(Equal("invalid map size: hello"))
		})
	})

	When("input correct format", func() {
		It("should return error = nil and add states to logger", func() {
			scanner := bufio.NewScanner(strings.NewReader("19\nF\nL"))

			states := make([]string, 0)

			err := ExecuteRoverWithScanner(scanner, func(s string) {
				states = append(states, s)
			})
			Expect(err).To(BeNil())
			Expect(states).To(Equal([]string{"N:0,0", "N:0,1", "W:0,1"}))
		})
	})

	When("invalid format", func() {
		It("should return error unsupport command", func() {
			scanner := bufio.NewScanner(strings.NewReader("24\nH"))

			err := ExecuteRoverWithScanner(scanner, func(_ string) {})
			Expect(err.Error()).To(Equal("unsupport command: H"))
		})
	})
})
