package service

type Provider struct {
	ProductService
}

func NewProvider() *Provider {
	return &Provider{
		ProductService: NewProductService(),
	}
}