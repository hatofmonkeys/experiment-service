package main_test

import (
	"net/url"

	. "github.com/hatofmonkeys/experiment-service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Experiment", func() {

	Describe("Specimen", func() {

		Context("when the ratio is set to 100% app A", func() {
			It("returns the provided url with '-a' appended to the hostname", func() {
				url, _ := url.Parse("https://host.google.com/")
				expectedUrl, _ := url.Parse("https://host-a.google.com/")

				ratio := 100

				modifiedUrl, _ := Specimen(url, ratio)

				Expect(modifiedUrl).To(Equal(expectedUrl))
			})
		})

		Context("when the ratio is set to 0% app A", func() {
			It("returns the provided url with '-b' appended to the hostname", func() {
				url, _ := url.Parse("https://host.google.com/")
				expectedUrl, _ := url.Parse("https://host-b.google.com/")

				ratio := 0

				modifiedUrl, _ := Specimen(url, ratio)

				Expect(modifiedUrl).To(Equal(expectedUrl))
			})
		})

		Context("when the ratio is set to 50% app A", func() {
			It("returns the provided url with either '-a' or '-b' appended to the hostname", func() {
				url, _ := url.Parse("https://host.google.com/")

				ratio := 50

				modifiedUrl, _ := Specimen(url, ratio)

				Expect(modifiedUrl.String()).To(MatchRegexp("https://host-[ab].google.com/"))
			})
		})

	})
})
