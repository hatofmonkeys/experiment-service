package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExperimentService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Experiment Suite")
}
