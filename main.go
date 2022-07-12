package main

import (
	"flag"
	"fmt"

	"github.com/zknill/java-graphviz/dag"
	"github.com/zknill/java-graphviz/parser"
)

func main() {
	var dir string
	var filter string
	var printOptions dag.DefaultPrintOptions

	flag.StringVar(&dir, "dir", "", "directory to parse")
	flag.StringVar(&filter, "filter", "", "package to filter by")
	flag.Var(&printOptions, "colorMap", "com.foo.bar.MyClass=red")

	flag.Parse()

	d := parser.DirParser(dir, filter)

	for _, line := range d.GraphViz(printOptions) {
		fmt.Println(line)
	}
}
