package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Home Handler", func() {
	var (
		handler http.Handler
		req     *http.Request
		w       *httptest.ResponseRecorder
		body    string
	)

	BeforeEach(func() {
		handler = Home()
		req = httptest.NewRequest("GET", "/", nil)
		w = httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		body = w.Body.String()
	})

	Context("when serving home page", func() {
		It("should return HTTP 200", func() {
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("should return HTML content type", func() {
			contentType := w.Header().Get("Content-Type")
			Expect(strings.Contains(contentType, "text/html")).To(BeTrue())
		})

		It("should include Hestia title", func() {
			Expect(body).To(ContainSubstring("Hestia"))
		})

		It("should include Pico CSS stylesheet", func() {
			Expect(body).To(ContainSubstring("pico.min.css"))
			Expect(body).To(ContainSubstring("cdn.jsdelivr.net"))
		})

		It("should display Spanish content", func() {
			Expect(body).To(ContainSubstring("Bienvenido"))
			Expect(body).To(ContainSubstring("Gestión Financiera"))
			Expect(body).To(ContainSubstring("Fecha actual"))
		})
	})
})