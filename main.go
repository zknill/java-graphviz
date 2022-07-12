package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zknill/java-graphviz/parser"
)

func main() {
	var file string

	flag.StringVar(&file, "file", "", "filename to parse")
	flag.Parse()

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d := parser.Parse(f)

	for _, line := range d.GraphViz() {
		fmt.Println(line)
	}
}
