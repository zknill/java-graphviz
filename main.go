package main

import (
	"flag"
	"fmt"

	"github.com/zknill/java-graphviz/parser"
)

func main() {
	var dir string
	var filter string

	flag.StringVar(&dir, "dir", "", "directory to parse")
	flag.StringVar(&filter, "filter", "", "package to filter by")
	flag.Parse()

	d := parser.DirParser(dir, filter)

	for _, line := range d.GraphViz() {
		fmt.Println(line)
	}
}
