package dag

import (
	"fmt"
	"strings"
)

type Node struct {
	Name     string
	Children []Node
	Package  bool
	Root     bool
}

func (n Node) GraphViz() []string {
	var out []string
	id := n.id()

	if n.Package {
		out = append(out, fmt.Sprintf("subgraph cluster_%s {", id))
		out = append(out, fmt.Sprintf("label=%q", n.Name))
		for _, c := range n.Children {
			out = append(out, fmt.Sprintf("%s[label=%q]", c.id(), c.Name))
		}

		out = append(out, "}")
	} else if !n.Root {
		out = append(out, fmt.Sprintf("%s[label=%q]", id, n.Name))
		for _, c := range n.Children {
			out = append(out, fmt.Sprintf("%s -> %s", id, c.id()))
		}
	}

	for _, c := range n.Children {
		out = append(out, c.GraphViz()...)
	}

	return out
}

func (n Node) id() string {
	return strings.ReplaceAll(n.Name, ".", "_")
}
