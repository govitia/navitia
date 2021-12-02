package types

type Route struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Frequent  bool   `json:"frequent"` // if the route has frequency or not. Can only be "False", but may be "True" in the future
	Line      Line   `json:"line"`
	Direction Place  `json:"direction"`
}
