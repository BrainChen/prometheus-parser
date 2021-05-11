package ast

import (
	"fmt"
	"strings"

	"github.com/BrainChen/prometheus-parser/models"
)

func AppendChild(ast *models.PrometheusAst, child models.AstChild) {
	ast.Dtype = "labeled"
	ast.Children = append(ast.Children, child)
}

func GetValue(ast models.PrometheusAst) string {
	return ast.Value
}

func SetValue(ast models.PrometheusAst, value string) models.PrometheusAst {
	ast.Value = value
	return ast
}

func GetCount(ast models.PrometheusAst) string {
	return ast.Count
}

func SetCount(ast models.PrometheusAst, count string) models.PrometheusAst {
	ast.Count = count
	return ast
}

func ReparseAst(ast models.PrometheusAst) string {
	var str strings.Builder
	switch ast.Dtype {
	case "comment":
		str.WriteString(ast.Value)
	case "labeled":
		str.WriteString(ast.Value)
		str.WriteString("{")
		fmt.Println(ast.Children)
		for i, v := range ast.Children {
			str.WriteString(v.TagName)
			str.WriteString("=")
			str.WriteString(v.Value)
			if i < len(ast.Children)-1 {
				str.WriteString(",")
			}
		}
		str.WriteString("}")
		str.WriteString(" ")
		str.WriteString(ast.Count)
	case "unlabeled":
		str.WriteString(ast.Value)
		str.WriteString(" ")
		str.WriteString(ast.Count)
	}
	return str.String()
}
