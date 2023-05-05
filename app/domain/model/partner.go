package model

type Partner struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Partners struct {
	Partners []Partner `json:"partners"`
}
