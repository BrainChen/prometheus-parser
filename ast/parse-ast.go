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
		// comment
		ast.Dtype = "comment"
		ast.Value = str
	} else if strings.Index(str, "{") >= 0 && strings.Index(str, "}") >= 0 {
		// labeled
		ast.Dtype = "labeled"
		var bf strings.Builder
		var count strings.Builder
		var parseName strings.Builder

		status := "preside"

		var child models.AstChild

		for _, v := range str {
			currentChar := string(v)
			if currentChar == " " {
				continue
			}
			switch status {
			case "preside":
				if currentChar == "{" {
					status = "inside"
					continue
				}
				bf.WriteString(currentChar)
			case "inside":
				switch currentChar {
				case "}":
					status = "outside"
					child.Value = parseName.String()
					parseName.Reset()
					ast.Children = append(ast.Children, child)
				case ",":
					child.Value = parseName.String()
					parseName.Reset()
					ast.Children = append(ast.Children, child)
				case "=":
					child.TagName = parseName.String()
					parseName.Reset()
				default:
					parseName.WriteString(currentChar)
				}
			case "outside":
				count.WriteString(currentChar)
			}
		}
		ast.Count = count.String()
		ast.Value = bf.String()
	} else {
		// unlabeled
		ast.Dtype = "unlabeled"
		var parseName strings.Builder
		for _, v := range str {
			currentChar := string(v)
			if currentChar == " " {
				ast.Value = parseName.String()
				parseName.Reset()
			} else {
				parseName.WriteString(currentChar)
			}
		}
		ast.Count = parseName.String()
	}
	return ast
}
