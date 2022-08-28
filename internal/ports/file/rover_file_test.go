package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	wd, _       = os.Getwd()
	projectRoot = filepath.Join(wd, "../../..")
)

var _ = Describe("Ports: file", func() {
	Describe("NewRoverFile", func() {
		It("correct file", func() {
			roverFile, err := NewRoverFile(fmt.Sprintf("%s/test.txt", projectRoot))
			Expect(err).To(BeNil())
			Expect(roverFile).To(BeAssignableToTypeOf(&RoverFile{}))
		})

		It("should return error when incorrect file", func() {
			roverFile, err := NewRoverFile(fmt.Sprintf("%s/test1.txt", projectRoot))
			Expect(err).NotTo(BeNil())
			Expect(roverFile).To(BeNil())
		})
	})

	Describe("Start", func() {
		It("should return log correct result", func() {
			roverFile, err := NewRoverFile(fmt.Sprintf("%s/test.txt", projectRoot))
			Expect(err).To(BeNil())

			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			roverFile.Start()

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = rescueStdout

			Expect(string(out)).To(Equal("N:0,0\nE:0,0\nE:1,0\nN:1,0\nN:1,1\nW:1,1\nS:1,1\nS:1,0\nW:1,0\n"))
		})
	})
})
