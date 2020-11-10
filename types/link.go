package types

type Link struct {
	Href      string `json:"href"`
	Type      string `json:"type"`
	Rel       string `json:"rel"`
	Templated bool   `json:"templated"`
}
