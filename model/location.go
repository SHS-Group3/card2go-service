package model

type Location struct {
	Name        string
	Description string
	Address     string
}

type Locatable interface {
	Location
}
