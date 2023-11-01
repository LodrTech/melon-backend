package model

type Product struct {
	ID 			string 	`jsonapi:"primary,products"`
	Name 		string 	`jsonapi:"attr,name"`
	Description string 	`jsonapi:"attr,description"`
	Price 		float64 `jsonapi:"attr,price"`
	Weight		float64 `jsonapi:"attr,weight"`
}