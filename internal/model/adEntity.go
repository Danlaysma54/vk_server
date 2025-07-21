package model

type AdEntity struct {
	AdId        string `json:"adId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Price       int    `json:"price"`
	Username    string `json:"username"`
}
type AdEntityForAuth struct {
	AdId        string `json:"adId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Price       int    `json:"price"`
	Username    string `json:"username"`
	Mine        bool   ` json:"mine"`
}
