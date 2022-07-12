package dag_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/zknill/java-graphviz/dag"
)

var _ = Describe("Graph", func() {
	It("prints the dag", func() {
		graph := dag.Node{
			Name: "parser.fixtures", Children: []dag.Node{
				{
					Name: "MyClass", Children: []dag.Node{
						{
							Name: "foo.bar.baz",
						},
					},
				},
			},
		}

		lines := graph.GraphViz()

		Expect(lines).To(Equal(
			[]string{
				"parser_fixtures[label=\"parser.fixtures\"]",
				"parser_fixtures -> MyClass",
				"MyClass[label=\"MyClass\"]",
				"MyClass -> foo_bar_baz",
				"foo_bar_baz[label=\"foo.bar.baz\"]",
			},
		))
	})
})
