package models

type Serie struct {
	Name     string  `json:"name"`
	Star     int     `json:"star"`
	Favorite bool	 `json:"favorite"`
}