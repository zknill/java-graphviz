package parser_test

import (
	"os"
	"sort"

	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zknill/java-graphviz/dag"
	"github.com/zknill/java-graphviz/parser"
)

var _ = Describe("Parser", func() {
	It("should parse a single class file", func() {
		file, err := os.Open("./fixtures/single/MyClass.java")
		Expect(err).NotTo(HaveOccurred())

		node := parser.Parse(file, "")
		Expect(node).To(
			Equal(
				dag.Node{
					Name:    "parser.fixtures.single",
					Package: true,
					Children: []dag.Node{
						{
							Name: "parser.fixtures.single.MyClass", Children: []dag.Node{
								{
									Name: "foo.bar.baz",
								},
							},
						},
					},
				},
			))
	})

	It("should parse a single enum file", func() {
		file, err := os.Open("./fixtures/enumeration/MyEnum.java")
		Expect(err).NotTo(HaveOccurred())

		node := parser.Parse(file, "")
		Expect(node).To(
			Equal(
				dag.Node{
					Name:    "parser.fixtures.enumeration",
					Package: true,
					Children: []dag.Node{
						{
							Name: "parser.fixtures.enumeration.MyEnum",
						},
					},
				},
			))
	})

	It("should parse a single record file", func() {
		file, err := os.Open("./fixtures/arecord/MyRecord.java")
		Expect(err).NotTo(HaveOccurred())

		node := parser.Parse(file, "")
		Expect(node).To(
			Equal(
				dag.Node{
					Name:    "parser.fixtures.arecord",
					Package: true,
					Children: []dag.Node{
						{
							Name: "parser.fixtures.arecord.MyRecord",
						},
					},
				},
			))
	})

	It("should parse a single interface file", func() {
		file, err := os.Open("./fixtures/iface/MyInterface.java")
		Expect(err).NotTo(HaveOccurred())

		node := parser.Parse(file, "")
		Expect(node).To(
			Equal(
				dag.Node{
					Name:    "parser.fixtures.iface",
					Package: true,
					Children: []dag.Node{
						{
							Name: "parser.fixtures.iface.MyInterface",
						},
					},
				},
			))
	})

	It("should parse a directory of files", func() {
		d := parser.DirParser("./fixtures/multi", "")
		want :=
			dag.Node{
				Name: "root",
				Children: []dag.Node{
					{
						Name: "parser.fixtures.multi",
						Children: []dag.Node{
							{
								Name: "parser.fixtures.multi.MyRoot",
								Children: []dag.Node{
									{
										Name:     "parser.fixtures.multi.nested.BirdsNest",
										Children: nil,
										Package:  false,
									},
									{
										Name:     "parser.fixtures.multi.nested.Nested",
										Children: nil,
										Package:  false,
									},
								},
								Package: false,
							},
						},
						Package: true,
					},
					{
						Name: "parser.fixtures.multi.nested",
						Children: []dag.Node{
							{
								Name:     "parser.fixtures.multi.nested.BirdsNest",
								Children: nil,
								Package:  false,
							},
							{
								Name:     "parser.fixtures.multi.nested.Nested",
								Children: nil,
								Package:  false,
							},
						},
						Package: true,
					},
				},
				Package: false,
			}

		sort.Slice(d.Children, func(i, j int) bool {
			return d.Children[i].Name < d.Children[j].Name
		})

		diff := cmp.Diff(d, want)

		Expect(diff).To(BeEmpty())

	})
})
