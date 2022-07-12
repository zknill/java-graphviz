package parser_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zknill/java-graphviz/dag"
	"github.com/zknill/java-graphviz/parser"
)

var _ = Describe("Parser", func() {

	It("should return correct nodes", func() {

		file, err := os.Open("./fixtures/MyClass.java")
		Expect(err).NotTo(HaveOccurred())

		node := parser.Parse(file)
		Expect(node).To(
			Equal(
				dag.Node{
					Name: "parser.fixtures", Children: []dag.Node{
						{
							Name: "MyClass", Children: []dag.Node{
								{
									Name: "foo.bar.baz",
								},
							},
						},
					},
				},
			))

	})
})
