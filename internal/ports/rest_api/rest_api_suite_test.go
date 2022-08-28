package rest_api

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegrationTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RestAPI Test Suite")
}
