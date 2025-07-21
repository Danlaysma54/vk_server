package model

type AdEntity struct {
	AdId        string
	Name        string
	Description string
	ImageUrl    string
	Price       int
	Username    string
}
type AdEntityForAuth struct {
	AdId        string
	Name        string
	Description string
	ImageUrl    string
	Price       int
	Username    string
	Mine        bool
}
