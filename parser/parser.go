package parser

import (
	"bufio"
	"io"
	"regexp"
	"strings"

	"github.com/zknill/java-graphviz/dag"
)

func Parse(src io.Reader) dag.Node {
	sc := bufio.NewScanner(src)

	sc.Split(bufio.ScanLines)

	pkgPattern := regexp.MustCompile("\\s*package\\s+([.?\\w]+)\\s*;")
	importPattern := regexp.MustCompile("\\s*import\\s+([.?\\w]+)\\s*;")
	classPattern := regexp.MustCompile("\\s*(?:public|private|protected)?\\s+(?:final|abstract)?\\s*(?:class|interface)\\s+(\\w+).*{.*")

	var imports []dag.Node
	var pkgNode dag.Node
	var classNode dag.Node

	for sc.Scan() {
		line := sc.Text()

		if matches := pkgPattern.FindStringSubmatch(line); len(matches) > 1 {
			pkgNode = dag.Node{Name: matches[1]}
		}

		if matches := importPattern.FindStringSubmatch(line); len(matches) > 1 {
			if strings.HasPrefix(matches[1], "com.askattest.site.model") {
				continue
			}
			if !strings.HasPrefix(matches[1], "com.askattest") {
				continue
			}
			imports = append(imports, dag.Node{Name: matches[1]})
		}

		if matches := classPattern.FindStringSubmatch(line); len(matches) > 1 {
			classNode = dag.Node{Name: matches[1], Children: imports}
			break
		}
	}

	pkgNode.Children = []dag.Node{classNode}

	return pkgNode
}
