package models

type PrometheusAst struct {
	Dtype    string     `json:"dtype"`
	Value    string     `json:"value"`
	Children []AstChild `json:"ast_child"`
	Count    string     `json:"count"`
}

type AstChild struct {
	TagName string `json:"tag_name"`
	Value   string `json:"value"`
}
