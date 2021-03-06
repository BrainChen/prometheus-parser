package ast

import (
	"strings"

	"github.com/BrainChen/prometheus-parser/models"
)

// Dtype=comment # TYPE rebuffer_reason counter
// Dtype=labeled rebuffer_reason{name="high-cpu-usage",label="argus"} 155
// Dtype=unlabeled gc_pause_total 498956470214
func ParseAst(str string) models.PrometheusAst {
	var ast models.PrometheusAst
	if strings.HasPrefix(str, "#") {
		ast.Dtype = "comment" // comment type
		ast.Value = str
	} else if strings.Contains(str, "{") && strings.Contains(str, "}") {
		ast.Dtype = "labeled"
		var bf strings.Builder
		var count strings.Builder
		var parseName strings.Builder
		status := "preside"
		var child models.AstChild
		for _, v := range str {
			if v == ' ' {
				continue
			}
			switch status {
			case "preside":
				if v == '{' {
					status = "inside"
					continue
				}
				bf.WriteRune(v)

			case "inside":
				switch v {
				case '}':
					status = "outside"
					child.Value = parseName.String()
					parseName.Reset()
					ast.Children = append(ast.Children, child)
				case ',':
					child.Value = parseName.String()
					parseName.Reset()
					ast.Children = append(ast.Children, child)
				case '=':
					child.TagName = parseName.String()
					parseName.Reset()
				default:
					parseName.WriteRune(v)
				}
			case "outside":
				count.WriteRune(v)
			}
		}
		ast.Count = count.String()
		ast.Value = bf.String()
	} else {
		ast.Dtype = "unlabeled"
		var parseName strings.Builder
		for _, v := range str {
			if v == ' ' {
				ast.Value = parseName.String()
				parseName.Reset()
			} else {
				parseName.WriteRune(v)
			}
		}
		ast.Count = parseName.String()
	}
	return ast
}
