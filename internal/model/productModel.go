package model

type Product struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
