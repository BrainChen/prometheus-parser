package main

import (
	"fmt"

	"github.com/BrainChen/prometheus-parser/ast"
	"github.com/BrainChen/prometheus-parser/models"
)

func main() {
	a := ast.ParseAst(`gc_pause_total 498956470214`)
	fmt.Println(a.Count, a.Dtype, a.Value)
	var child models.AstChild
	child.TagName = "looooooo"
	child.Value = "12321312"
	ast.AppendChild(&a, child)
	b := ast.ReparseAst(a)
	fmt.Println(b)
}
