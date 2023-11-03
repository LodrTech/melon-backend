package model

type Product struct {
	ID 			string 	`jsonapi:"primary,products"`
	Name 		string 	`jsonapi:"attr,name"`
	Description string 	`jsonapi:"attr,description"`
	Price 		float32 `jsonapi:"attr,price"`
	Weight		int 	`jsonapi:"attr,weight"`
}