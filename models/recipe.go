package models

type Recipe struct {
	Id       uint   `json:"recipeId"`
	Title    string `json:"title"`
	Describ  string `json:"describ"`
	TimePrep uint   `json:"timePrep"`
	Image    string `json:"image"`
	Author   string `json:"author"`
	Category string `json:"category"`
}
