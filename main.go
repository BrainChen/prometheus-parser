package main

import (
	"fmt"
	"strings"

	"github.com/BrainChen/prometheus-parser/ast"
)

func main() {
	result := "gc_pause_total 498956470214\ngc_pause_total 498956470214\n# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.\ngo_info{version=\"go1.16.2\"} 1"
	temp := strings.Split(result, "\n")
	for _, v := range temp {
		a := ast.ParseAst(v)
		ast.AppendLabel(&a, "looooooo", "1232311")
		b := ast.ReparseAst(a)
		fmt.Println(b)
	}

}
