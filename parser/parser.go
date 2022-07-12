package parser

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/zknill/java-graphviz/dag"
)

func Parse(src io.Reader, filter string) dag.Node {
	sc := bufio.NewScanner(src)

	sc.Split(bufio.ScanLines)

	pkgPattern := regexp.MustCompile("\\s*package\\s+([.?\\w]+)\\s*;")
	importPattern := regexp.MustCompile("\\s*import\\s+([.?\\w]+)\\s*;")
	classPattern := regexp.MustCompile(`\s*(?:public|private|protected)?\s*(?:final|abstract)?\s*(?:class|record|enum|interface)\s+(\w+).*{.*`)

	var imports []dag.Node
	var pkgNode dag.Node
	var classNode dag.Node

	for sc.Scan() {
		line := sc.Text()

		if matches := pkgPattern.FindStringSubmatch(line); len(matches) > 1 {
			pkgNode = dag.Node{Name: matches[1], Package: true}
		}

		if matches := importPattern.FindStringSubmatch(line); len(matches) > 1 {
			if filterOut(filter, matches[1]) {
				continue
			}

			imports = append(imports, dag.Node{Name: matches[1]})
		}

		if matches := classPattern.FindStringSubmatch(line); len(matches) > 1 {
			classNode = dag.Node{
				Name:     pkgNode.Name + "." + matches[1],
				Children: imports,
			}
			break
		}
	}

	if classNode.Name != "" {
		pkgNode.Children = []dag.Node{classNode}
	}

	return pkgNode
}

func filterOut(filter string, node string) bool {
	return filter != "" && !strings.HasPrefix(node, filter)
}

func DirParser(dirPath string, pkgFilter string) dag.Node {
	files, _ := ioutil.ReadDir(dirPath)
	var nodes []dag.Node
	for _, file := range files {
		fName := path.Join(dirPath, file.Name())

		if file.IsDir() {
			d := DirParser(fName, pkgFilter)
			nodes = append(nodes, d)
		}

		if !strings.HasSuffix(file.Name(), ".java") {
			continue
		}

		f, err := os.Open(fName)
		if err != nil {
			fmt.Println(err)

		}

		nodes = append(nodes, Parse(f, pkgFilter))
	}

	packages := make(map[string]dag.Node)

	for _, n := range nodes {
		if !n.Package {
			continue
		}

		cached, found := packages[n.Name]
		if !found {
			packages[n.Name] = n
			continue
		}

		cached.Children = append(cached.Children, n.Children...)

		packages[n.Name] = cached
	}

	packageNodes := make([]dag.Node, 0, len(packages))

	for _, pn := range packages {
		packageNodes = append(packageNodes, pn)
	}

	if len(packageNodes) == 1 {
		return packageNodes[0]
	}

	return dag.Node{
		Name:     "root",
		Children: packageNodes,
	}
}
