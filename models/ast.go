package models

type PrometheusAst struct {
	Dtype    string
	Value    string
	Children []AstChild
	Count    string
}

type AstChild struct {
	TagName string
	Value   string
}
