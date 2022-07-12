package dag

import (
	"fmt"
	"strings"
)

const defaultColor = "default"

type DefaultPrintOptions struct {
	m map[string]string
}

func (d DefaultPrintOptions) Color(name string) string {
	if len(d.m) == 0 {
		return defaultColor
	}

	color, found := d.m[name]
	if !found {
		return defaultColor
	}

	return color
}

func (d *DefaultPrintOptions) Set(val string) error {
	if d.m == nil {
		d.m = map[string]string{}
	}

	split := strings.Split(val, "=")
	d.m[split[0]] = split[1]

	return nil
}

func (d *DefaultPrintOptions) String() string {
	return fmt.Sprintf("%s", d.m)
}
