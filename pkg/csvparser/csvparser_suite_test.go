package csvparser_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCsvparser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Csvparser Suite")
}
