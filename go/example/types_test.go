package example

import (
	"github.com/golang/mock/gomock"
	g "github.com/onsi/ginkgo"
	m "github.com/onsi/gomega"
	"github.com/yangzuo0621/codebase/go/example/mock_example"
)

var _ = g.Describe("", func() {
	mockCtl := gomock.NewController(g.GinkgoT())
	g.It("", func() {
		foo := mock_example.NewMockFoo(mockCtl)
		foo.EXPECT().Bar().Return("bar", nil)

		result, err := foo.Bar()
		m.Expect(err).NotTo(m.HaveOccurred())
		m.Expect(result).To(m.Equal("bar"))
	})
})
