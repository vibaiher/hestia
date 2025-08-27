package static

import (
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Static", func() {
	Context("when accessing static files", func() {
		It("should find existing template file", func() {
			path := File("templates/index.html")
			_, err := os.Stat(path)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should find existing locale file", func() {
			path := File("locales/es.json")
			_, err := os.Stat(path)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return absolute path", func() {
			path := File("templates/index.html")
			Expect(filepath.IsAbs(path)).To(BeTrue())
		})

		It("should contain static directory in path", func() {
			path := File("templates/index.html")
			Expect(strings.Contains(path, "static")).To(BeTrue())
		})
	})
})