package example

import (
	"testing"

	g "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	m "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	m.RegisterFailHandler(g.Fail)
	g.RunSpecsWithDefaultAndCustomReporters(
		t,
		"example test suites",
		[]g.Reporter{
			reporters.NewJUnitReporter("junit.xml"),
		},
	)
}
