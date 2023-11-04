package model

type Product struct {
	ID 			string 	`jsonapi:"primary,products"`
	Name 		string 	`jsonapi:"attr,name" validate:"required"`
	Description string 	`jsonapi:"attr,description"`
	Price 		float32 `jsonapi:"attr,price" validate:"required,gt=0"`
	Weight		int 	`jsonapi:"attr,weight" validate:"required"`
}

type ProductListRequest struct {

}

type ProductUpdateRequest struct {
	ID 			string 	`jsonapi:"primary,products"`
	Name 		string 	`jsonapi:"attr,name"`
	Description string 	`jsonapi:"attr,description"`
	Price 		float32 `jsonapi:"attr,price" validate:"gt=0"`
	Weight		int 	`jsonapi:"attr,weight" validate:"gt=0"`
}