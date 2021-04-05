package tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAutomationPracticeWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Registration test suite")
}
