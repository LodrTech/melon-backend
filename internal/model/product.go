package model

type Product struct {
	ID 			string 	`jsonapi:"primary,products"`
	Name 		string 	`jsonapi:"attr,name" validate:"required"`
	Description string 	`jsonapi:"attr,description"`
	Price 		float32 `jsonapi:"attr,price" validate:"required,gt=0"`
	Weight		int 	`jsonapi:"attr,weight" validate:"required"`
}