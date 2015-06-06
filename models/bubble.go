package models

type Bubble struct {
	Name     string  `json:"name"`
	Size     int     `json:"size"`
	Children []Bubble `json:"children"`
}

