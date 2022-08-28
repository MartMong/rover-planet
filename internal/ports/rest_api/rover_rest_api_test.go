package rest_api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ports: REST_API", func() {
	Describe("NewRoverRestAPI", func() {
		It("should return RoverRestAPI", func() {
			roverRestAPI := NewRoverRestAPI(4000)
			Expect(roverRestAPI).To(BeAssignableToTypeOf(&RoverRestAPI{}))
		})
	})

	Describe("handlerRover", func() {
		It("should execute and write result to ResponseWriter", func() {
			reader := strings.NewReader("24\nR\nF\nL\nF\nL\nL\nF\nR")

			request, err := http.NewRequest(http.MethodPost, "http://localhost:8080", reader)
			Expect(err).To(BeNil())

			w := httptest.NewRecorder()

			handlerRover(w, request)

			fmt.Println(w.Body.String())

			Expect(w.Body.String()).To(Equal("N:0,0\nE:0,0\nE:1,0\nN:1,0\nN:1,1\nW:1,1\nS:1,1\nS:1,0\nW:1,0"))
		})
	})
})
