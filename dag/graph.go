package dag

import (
	"fmt"
	"strings"
)

type Node struct {
	Name     string
	Children []Node
	Package  bool
}

func (n Node) GraphViz() []string {
	var out []string
	id := n.id()

	out = append(out, fmt.Sprintf("%s[label=%q]", id, n.Name))

	for _, c := range n.Children {
		out = append(out, fmt.Sprintf("%s -> %s", id, c.id()))

		out = append(out, c.GraphViz()...)
	}

	return out
}

func (n Node) id() string {
	return strings.ReplaceAll(n.Name, ".", "_")
}
