package ast

import (
	"strings"

	"github.com/BrainChen/prometheus-parser/models"
)

func AppendLabel(ast *models.PrometheusAst, label, value string) {
	modified_value := "\"" + value + "\""
	var child models.AstChild
	child.TagName = label
	child.Value = modified_value
	AppendChild(ast, child)

}
func AppendChild(ast *models.PrometheusAst, child models.AstChild) {
	if ast.Dtype != "comment" {
		ast.Dtype = "labeled"
		ast.Children = append(ast.Children, child)
	}
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
	var strBuilder strings.Builder
	switch ast.Dtype {
	case "comment":
		strBuilder.WriteString(ast.Value)
		strBuilder.WriteRune('\n')
	case "labeled":
		strBuilder.WriteString(ast.Value)
		strBuilder.WriteRune(' ')
		strBuilder.WriteRune('{')
		for i, v := range ast.Children {
			strBuilder.WriteString(v.TagName)
			strBuilder.WriteRune('=')
			strBuilder.WriteString(v.Value)
			if i < len(ast.Children)-1 {
				strBuilder.WriteRune(',')
			}
		}
		strBuilder.WriteRune('}')
		strBuilder.WriteRune(' ')
		strBuilder.WriteString(ast.Count)
		strBuilder.WriteRune('\n')
	case "unlabeled":
		strBuilder.WriteString(ast.Value)
		strBuilder.WriteRune(' ')
		strBuilder.WriteString(ast.Count)
		strBuilder.WriteRune('\n')
	}
	return strBuilder.String()
}
