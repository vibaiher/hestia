package i18n

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Translations", func() {
	Context("when loading Spanish translations", func() {
		It("should load translations successfully", func() {
			translations, err := Load()
			Expect(err).ToNot(HaveOccurred())
			Expect(translations.App.Name).To(Equal("Hestia"))
			Expect(translations.Home.Welcome).To(Equal("Bienvenido a Hestia"))
			Expect(translations.Features.Title).To(Equal("Próximas funcionalidades"))
		})

		It("should have all required sections", func() {
			translations, err := Load()
			Expect(err).ToNot(HaveOccurred())
			
			Expect(translations.App.Name).ToNot(BeEmpty())
			Expect(translations.Home.Welcome).ToNot(BeEmpty())
			Expect(translations.Features.Title).ToNot(BeEmpty())
			Expect(translations.Footer.Tagline).ToNot(BeEmpty())
		})
	})
})