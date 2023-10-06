package models

type Product struct {
	ID    uint   `json:"id"`
	Title string `json: "title"`
	Desc  string `json:"desc"`
}
